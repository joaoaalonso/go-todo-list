package routes

import (
	"github.com/gorilla/mux"
	"github.com/joaoaalonso/go-todo-list/interfaces/http/controllers"
)

func userRoutes(router *mux.Router) {
	r := router.PathPrefix("/user").Subrouter()

	userController := controllers.UserController{}

	r.HandleFunc("", userController.Get).Methods("GET")
	r.HandleFunc("", userController.Update).Methods("PUT")
	r.HandleFunc("", userController.Create).Methods("POST")
	r.HandleFunc("/login", userController.Login).Methods("POST")
}
