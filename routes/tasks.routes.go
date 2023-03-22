package routes

import (
	"encoding/json"
	"github.com/diegoamc92/go-gorm-restapi/db"
	"github.com/diegoamc92/go-gorm-restapi/models"
	"github.com/gorilla/mux"
	"net/http"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Tasks
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Tasks
	json.NewDecoder(r.Body).Decode(&task)
	createTask := db.DB.Create(&task)
	err := createTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&task)
}
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Tasks
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}
	json.NewEncoder(w).Encode(&task)
}
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Tasks
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}
	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusNoContent)
}
