package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ivanleodomin/go-api-tasks/database"
	"github.com/ivanleodomin/go-api-tasks/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	fmt.Println(params["id"])
	database.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)

		m := make(map[string]string)
		m["error"] = "User not found"
		json.NewEncoder(w).Encode(m)
		return
	}

	json.NewEncoder(w).Encode(&user)

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
	var user models.User
	params := mux.Vars(r)
	fmt.Println(params["id"])
	database.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)

		m := make(map[string]string)
		m["error"] = "User not found"
		json.NewEncoder(w).Encode(m)
		return
	}

	database.DB.Unscoped().Delete(&user)
	json.NewEncoder(w).Encode(&user)
}
func UpateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get users"))

}
