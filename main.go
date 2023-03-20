package main

import (
	"net/http"

	"github.com/diegoamc92/go-gorm-restapi/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.Home)

	http.ListenAndServe(":3000", r)
}
