package controllers

import (
	"fmt"
	"net/http"
)

// TodoController handle all router from /todo
type TodoController struct{}

// List all todo
func (todoController *TodoController) List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "todo controller list")
}

// Get one todo
func (todoController *TodoController) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "todo controller get")
}

// Create a new todo
func (todoController *TodoController) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "todo controller create")
}

// Update a todo
func (todoController *TodoController) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "todo controller update")
}

// Delete a todo
func (todoController *TodoController) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "todo controller delete")
}
