package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"./db"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("PORT")

	fmt.Println("Server running in port " + port)

	db.DB = db.SetupDB()
	defer db.DB.Close()

	log.Fatal(http.ListenAndServe(":"+port, routerConfig()))
}
