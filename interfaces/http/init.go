package http

import (
	"log"
	"net/http"
	"os"

	"github.com/joaoaalonso/go-todo-list/interfaces/http/middlewares"
	"github.com/joaoaalonso/go-todo-list/interfaces/http/routes"

	"github.com/gorilla/mux"
)

// Init register all routes from the http interface and start the server
func Init() {
	router := mux.NewRouter()
	routes.Register(router)

	router.Use(middlewares.LoggingHandler)
	router.Use(middlewares.CompressHandler)
	router.Use(middlewares.CorsHandler)

	port := os.Getenv("DB_HOST")
	if port == "" {
		port = "8000"
	}

	log.Println("Server is running on port " + port)
	http.ListenAndServe(":"+port, router)
}
