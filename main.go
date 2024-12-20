package main

import (
	"net/http"
	"os"
	"rest-api-go/database"
	"rest-api-go/handlers"
)

func main() {
	database.InitializeDatabase()
	defer database.DB.Close()
	handlers.UserHandler()
	port := os.Getenv("PORT")
	if port == "" {
		port = "6001"
	}
	http.ListenAndServe(":"+port, nil)
}
