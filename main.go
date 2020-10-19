package main

import (
	"fmt"
	"log"

	"github.com/joaoaalonso/go-todo-list/interfaces/http"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.Init()

	fmt.Println("ae")
}
