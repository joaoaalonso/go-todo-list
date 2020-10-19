package routes

import (
	"github.com/gorilla/mux"
	"github.com/joaoaalonso/go-todo-list/interfaces/http/controllers"
)

func todoRoutes(router *mux.Router) {
	r := router.PathPrefix("/todo").Subrouter()

	todoController := controllers.TodoController{}

	r.HandleFunc("", todoController.List).Methods("GET")
	r.HandleFunc("", todoController.Create).Methods("POST")
	r.HandleFunc("/{id}", todoController.Get).Methods("GET")
	r.HandleFunc("/{id}", todoController.Update).Methods("PUT")
	r.HandleFunc("/{id}", todoController.Delete).Methods("DELETE")
}
