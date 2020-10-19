package controllers

import (
	"fmt"
	"net/http"
)

// UserController handle all router from /user
type UserController struct{}

// Get return current user data
func (userController *UserController) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "user controller get")
}

// Update user data
func (userController *UserController) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "user controller update")
}

// Create new user
func (userController *UserController) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "user controller create")
}

// Login authenticate user from email and password and return a token
func (userController *UserController) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "user controller login")
}
