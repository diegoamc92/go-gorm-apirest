package main

import (
	"net/http"

	"github.com/diegoamc92/go-gorm-restapi/db"
	"github.com/diegoamc92/go-gorm-restapi/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()
	r := mux.NewRouter()
	r.HandleFunc("/", routes.Home)

	http.ListenAndServe(":3000", r)
}
