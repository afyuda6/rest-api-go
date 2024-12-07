package main

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-api-go/database"
	"rest-api-go/handlers"
)

type Response struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data,omitempty"`
	Errors []string    `json:"errors,omitempty"`
}

func main() {
	database.InitializeDatabase()
	defer database.DB.Close()

	http.HandleFunc("/users", handlers.UserHandler)
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/users" || r.URL.Path == "/users/" {
			handlers.UserHandler(w, r)
		} else {
			response := Response{
				Status: "Not Found",
				Code:   404,
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users" && r.URL.Path != "/users/" {
			response := Response{
				Status: "Not Found",
				Code:   404,
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
			return
		}
	})

	if err := http.ListenAndServe(":6001", nil); err != nil {
		log.Fatalf("Server failed!\n")
	}
}
