package main

import (
	"log"
	"net/http"

	"rest-api-go/database"
	"rest-api-go/handlers"
)

func main() {
	database.InitializeDatabase()
	defer database.DB.Close()

	http.HandleFunc("/users", handlers.UserHandler)

	log.Println("Server running on port 6001")
	if err := http.ListenAndServe(":6001", nil); err != nil {
		log.Fatalf("Server failed to start: %s", err)
	}
}
