package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"time"

	"../db"
	"../helpers"
	"../middlewares"
	"../structs"

	"github.com/badoux/checkmail"
	fStructs "github.com/fatih/structs"
)

//UsersLogin user login handler
func UsersLogin(w http.ResponseWriter, r *http.Request) {
	type error interface {
		Error() string
	}
	var (
		usersStruct    structs.Users
		responseStruct structs.ResponseStruct
	)
	w.Header().Set("Content-Type", "application/json")
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		responseStruct.Status = false
		responseStruct.Message = err.Error()
		responseStruct.Result = nil
		json.NewEncoder(w).Encode(&responseStruct)
		return
	}
	if data["username"] == "" || data["username"] == nil || data["password"] == "" || data["password"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		responseStruct.Status = false
		responseStruct.Message = "Login gagal"
		responseStruct.Result = nil
		json.NewEncoder(w).Encode(&responseStruct)
		return
	}
	if err := db.DB.Where(map[string]interface{}{"username": data["username"]}).Find(&usersStruct).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		responseStruct.Status = false
		responseStruct.Message = "Login gagal"
		responseStruct.Result = nil
	} else {
		hashPassword := usersStruct.Password
		usersStruct.Password = fmt.Sprintf("%v", data["password"])
		if hashPassword == helpers.GeneratePassword(usersStruct) {
			token, err := middlewares.GenerateJWT(usersStruct)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				responseStruct.Status = false
				responseStruct.Message = "Login gagal"
				responseStruct.Result = nil
			} else {
				responseStruct.Status = true
				responseStruct.Message = "Login Berhasil"
				responseStruct.Result = map[string]interface{}{
					"token": token,
				}
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			responseStruct.Status = false
			responseStruct.Message = "Login gagal"
			responseStruct.Result = nil
		}
	}

	json.NewEncoder(w).Encode(&responseStruct)
}

//UsersRegisterHandler user register handler
func UsersRegisterHandler(w http.ResponseWriter, r *http.Request) {
	type error interface {
		Error() string
	}
	var (
		usersStruct    structs.Users
		responseStruct structs.ResponseStruct
	)
	w.Header().Set("Content-Type", "application/json")
	decodeRequest := json.NewDecoder(r.Body)
	currentDate := time.Now()
	if errDecode := decodeRequest.Decode(&usersStruct); errDecode != nil {
		w.WriteHeader(http.StatusBadRequest)
		responseStruct.Status = false
		responseStruct.Message = errDecode.Error()
		responseStruct.Result = nil
		json.NewEncoder(w).Encode(&responseStruct)
		return
	}
	usersStruct.CreatedAt = currentDate
	usersStruct.UpdatedAt = currentDate

	reflectedValue := reflect.ValueOf(usersStruct)
	reflectedField := reflect.TypeOf(usersStruct)
	lengthField := reflectedField.NumField()

	for i := 0; i < lengthField; i++ {
		value := reflectedValue.Field(i).String()
		valid := true
		message := ""
		if field := reflectedField.Field(i).Name; field == "Username" || field == "Email" || field == "Password" {
			if field == "Email" {
				if err := checkmail.ValidateFormat(value); err != nil {
					valid = false
					message = "Email tidak valid"
				}
			} else if field == "Username" {
				if 8 > len(value) || len(value) > 32 {
					valid = false
					message = "Username tidak boleh kurang dari 8 dan lebih dari 32 karakter"
				} else if matched, _ := regexp.MatchString(`^[^a-z]|[^a-z0-9]`, value); matched {
					valid = false
					message = "Username tidak valid, hanya boleh terdiri dari huruf kecil atau angka, dan diawal dengan huruf"
				} else if err := db.DB.Where(map[string]interface{}{field: value}).Find(&usersStruct).Error; err == nil {
					message = reflectedField.Field(i).Name + " Telah Digunakan"
					valid = false
				}
			} else if field == "Password" {
				if 8 > len(value) || len(value) > 32 {
					valid = false
					message = "Password tidak boleh kurang dari 8 dan lebih dari 32 karakter"
				}
			}
		}
		if value == "" {
			valid = false
			message = reflectedField.Field(i).Name + " Kosong"
		}
		if valid != true {
			w.WriteHeader(http.StatusBadRequest)
			responseStruct.Status = false
			responseStruct.Message = message
			responseStruct.Result = nil
			json.NewEncoder(w).Encode(&responseStruct)
			return
		}
	}

	usersStruct.ID = 0
	usersStruct.Password = helpers.GeneratePassword(usersStruct)

	if err := db.DB.Create(&usersStruct).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		responseStruct.Status = false
		responseStruct.Message = err.Error()
		responseStruct.Result = nil
	} else {
		responseStruct.Status = true
		responseStruct.Message = "Registrasi Berhasil"
		responseStruct.Result = fStructs.Map(&usersStruct)
	}
	json.NewEncoder(w).Encode(&responseStruct)
}
