package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ivanleodomin/go-api-tasks/database"
	"github.com/ivanleodomin/go-api-tasks/models"
	"github.com/ivanleodomin/go-api-tasks/routes"
)

func main() {

	// Database Conection
	database.DBConnection()
	database.DB.AutoMigrate(models.Tasks{})
	database.DB.AutoMigrate(models.User{})

	// Mux Configuration
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.GetUserByIdHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.UpateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
