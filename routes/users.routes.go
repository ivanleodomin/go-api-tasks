package routes

import (
	"encoding/json"
	"net/http"

	"github.com/ivanleodomin/go-api-tasks/database"
	"github.com/ivanleodomin/go-api-tasks/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get users"))

}
func PostUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	createdUser := database.DB.Create(&user)

	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)
}
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get users"))

}
func UpateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get users"))

}
