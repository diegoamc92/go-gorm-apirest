package routes

import (
	"encoding/json"
	"github.com/diegoamc92/go-gorm-restapi/db"
	"github.com/diegoamc92/go-gorm-restapi/models"
	"net/http"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User"))
}
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	createUser := db.DB.Create(&user)
	err := createUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)

}
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete"))
}
