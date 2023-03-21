package routes

import "net/http"

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Users"))
}
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User"))
}
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Post"))
}
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete"))
}
