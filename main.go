package main

import (
	"net/http"

	"github.com/diegoamc92/go-gorm-restapi/models"

	"github.com/diegoamc92/go-gorm-restapi/db"
	"github.com/diegoamc92/go-gorm-restapi/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBconnection()
	db.DB.AutoMigrate(models.Tasks{})

	err := db.DB.AutoMigrate(models.User{})
	if err != nil {
		return
	}

	r := mux.NewRouter()

	//user route
	r.HandleFunc("/", routes.Home)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	//Tasks routes

	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
