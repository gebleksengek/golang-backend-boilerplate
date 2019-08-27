package main

import (
	"os"

	"./controllers"
	"./middlewares"
	"github.com/gorilla/mux"
)

func routerConfig() *mux.Router {
	pathPrefix := os.Getenv("API_PATH_PREFIX")
	router := mux.NewRouter().StrictSlash(true).PathPrefix(pathPrefix).Subrouter()

	router.Use(middlewares.RequestLogger)

	// Index Request
	router.HandleFunc("/", controllers.IndexHandlerGET).Methods("GET")

	// Users Request
	router.HandleFunc("/users/register", controllers.UsersRegisterHandler).Methods("POST")
	router.HandleFunc("/users/login", controllers.UsersLogin).Methods("POST")

	// Posts Request
	router.HandleFunc("/posts", controllers.ListPostHandler).Methods("GET")
	router.HandleFunc("/posts", middlewares.JWTMiddleware(controllers.CreatePostHandler)).Methods("POST")
	router.HandleFunc("/posts", middlewares.JWTMiddleware(controllers.DeletePostHandler)).Methods("DELETE")
	router.HandleFunc("/posts", middlewares.JWTMiddleware(controllers.UpdatePostHandler)).Methods("PATCH")

	return router
}
