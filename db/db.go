package db

import (
	"fmt"
	"os"

	"../structs"

	"github.com/jinzhu/gorm"
)

//DB global DB
var DB *gorm.DB

//SetupDB initialize database
func SetupDB() *gorm.DB {
	type error interface {
		Error() string
	}
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	name := os.Getenv("DATABASE_NAME")
	user := os.Getenv("DATABASE_LOGIN")
	pass := os.Getenv("DATABASE_PASS")

	db, err := gorm.Open("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+name+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	db.AutoMigrate(&structs.Users{}, &structs.Posts{})

	return db
}
