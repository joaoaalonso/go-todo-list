package routes

import "github.com/gorilla/mux"

// Register all routes to router
func Register(router *mux.Router) {
	todoRoutes(router)
	userRoutes(router)
}
