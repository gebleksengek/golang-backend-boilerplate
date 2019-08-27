package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"../db"
	"../structs"
	fStructs "github.com/fatih/structs"
)

//ListPostHandler List post handler
func ListPostHandler(w http.ResponseWriter, r *http.Request) {
	type error interface {
		Error() string
	}
	var (
		postsStruct             []structs.Posts
		responseMultiDataStruct structs.ResponseMultiDataStruct
	)
	w.Header().Set("Content-Type", "application/json")
	if err := db.DB.Find(&postsStruct).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		responseMultiDataStruct.Status = false
		responseMultiDataStruct.Message = "Gagal memuat konten"
		responseMultiDataStruct.Result = nil
	} else {
		var results = make([]map[string]interface{}, len(postsStruct))
		for i := 0; i < len(postsStruct); i++ {
			results[i] = fStructs.Map(postsStruct[i])
		}
		responseMultiDataStruct.Status = true
		responseMultiDataStruct.Message = "Berhasil membuat content"
		responseMultiDataStruct.Result = results
	}
	json.NewEncoder(w).Encode(&responseMultiDataStruct)
}

//CreatePostHandler create post handler
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	type error interface {
		Error() string
	}
	var (
		postsStruct    structs.Posts
		responseStruct structs.ResponseStruct
	)
	w.Header().Set("Content-Type", "application/json")
	currentDate := time.Now()
	if err := json.NewDecoder(r.Body).Decode(&postsStruct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		responseStruct.Status = false
		responseStruct.Message = err.Error()
		responseStruct.Result = nil
		json.NewEncoder(w).Encode(&responseStruct)
		return
	}
	if postsStruct.Title == "" ||
		postsStruct.Content == "" {
		w.WriteHeader(http.StatusBadRequest)
		responseStruct.Status = false
		responseStruct.Message = "Gagal membuat content"
		responseStruct.Result = nil
		json.NewEncoder(w).Encode(&responseStruct)
		return
	}
	postsStruct.CreatedAt = currentDate
	postsStruct.UpdatedAt = currentDate
	postsStruct.ID = 0
	if err := db.DB.Create(&postsStruct).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		responseStruct.Status = false
		responseStruct.Message = "Gagal membuat content"
		responseStruct.Result = nil
	} else {
		responseStruct.Status = true
		responseStruct.Message = "Berhasil membuat content"
		responseStruct.Result = fStructs.Map(&postsStruct)
	}
	json.NewEncoder(w).Encode(&responseStruct)
}

//DeletePostHandler Delete post handler
func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	type error interface {
		Error() string
	}
	var (
		postsStruct    structs.Posts
		responseStruct structs.ResponseStruct
	)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&postsStruct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		responseStruct.Status = false
		responseStruct.Message = "Gagal menerima data"
		responseStruct.Result = nil
		json.NewEncoder(w).Encode(&responseStruct)
		return
	}
	if postsStruct.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		responseStruct.Status = false
		responseStruct.Message = "Data tidak valid"
		responseStruct.Result = nil
		json.NewEncoder(w).Encode(&responseStruct)
		return
	}
	if err := db.DB.Where(map[string]interface{}{"id": postsStruct.ID}).Delete(&postsStruct).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		responseStruct.Status = false
		responseStruct.Message = "Gagal menghapus content"
		responseStruct.Result = nil
	} else {
		responseStruct.Status = true
		responseStruct.Message = "Berhasil menghapus content"
		responseStruct.Result = nil
	}
	json.NewEncoder(w).Encode(&responseStruct)
}

//UpdatePostHandler Update post handler
func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	type error interface {
		Error() string
	}
	var (
		postsStruct    structs.Posts
		responseStruct structs.ResponseStruct
	)
	w.Header().Set("Content-Type", "application/json")
	currentDate := time.Now()
	if err := json.NewDecoder(r.Body).Decode(&postsStruct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		responseStruct.Status = false
		responseStruct.Message = "Gagal menerima data"
		responseStruct.Result = nil
		json.NewEncoder(w).Encode(&responseStruct)
		return
	}
	if postsStruct.ID == 0 ||
		postsStruct.Title == "" ||
		postsStruct.Content == "" {
		w.WriteHeader(http.StatusBadRequest)
		responseStruct.Status = false
		responseStruct.Message = "Data tidak valid"
		responseStruct.Result = nil
		json.NewEncoder(w).Encode(&responseStruct)
		return
	}
	postsStruct.UpdatedAt = currentDate
	if err := db.DB.Model(&postsStruct).Updates(&postsStruct).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		responseStruct.Status = false
		responseStruct.Message = "Gagal mengubah content"
		responseStruct.Result = nil
	} else {
		responseStruct.Status = true
		responseStruct.Message = "Berhasil mengubah content"
		responseStruct.Result = fStructs.Map(&postsStruct)
	}
	json.NewEncoder(w).Encode(&responseStruct)
}
