package main

import (
	"net/http"
	"rest-api-go/database"
	"rest-api-go/handlers"
)

func main() {
	database.InitializeDatabase()
	defer database.DB.Close()
	handlers.UserHandler()
	http.ListenAndServe(":6001", nil)
}
