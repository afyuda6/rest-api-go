package handlers

import (
	"encoding/json"
	"net/http"
	"rest-api-go/database"
	"strconv"
	"strings"
)

type Response struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data,omitempty"`
	Errors string      `json:"errors,omitempty"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func UserHandler() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/users" || r.URL.Path == "/users/" {
			switch r.Method {
			case http.MethodGet:
				handleReadUsers(w)
			case http.MethodPost:
				handleCreateUser(w, r)
			case http.MethodPut:
				handleUpdateUser(w, r)
			case http.MethodDelete:
				handleDeleteUser(w, r)
			default:
				response := Response{
					Status: "Method Not Allowed",
					Code:   405,
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusMethodNotAllowed)
				json.NewEncoder(w).Encode(response)
			}
		} else {
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
}

func handleReadUsers(w http.ResponseWriter) {
	rows, _ := database.DB.Query("SELECT id, name FROM users")
	defer rows.Close()
	users := []User{}
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}
	response := Response{
		Status: "OK",
		Code:   200,
		Data:   users,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimSpace(r.FormValue("name"))
	if name == "" {
		errors := "Missing 'name' parameter"
		response := Response{
			Status: "Bad Request",
			Code:   400,
			Errors: errors,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	newUser := User{Name: name}
	database.DB.Exec("INSERT INTO users (name) VALUES (?)", newUser.Name)
	response := Response{
		Status: "Created",
		Code:   201,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimSpace(r.FormValue("name"))
	idStr := strings.TrimSpace(r.FormValue("id"))
	if idStr == "" || name == "" {
		errors := "Missing 'id' or 'name' parameter"
		response := Response{
			Status: "Bad Request",
			Code:   400,
			Errors: errors,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	id, _ := strconv.Atoi(idStr)
	newUser := User{ID: id, Name: name}
	database.DB.Exec("UPDATE users SET name = ? WHERE id = ?", newUser.Name, newUser.ID)
	response := Response{
		Status: "OK",
		Code:   200,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimSpace(r.FormValue("id"))
	if idStr == "" {
		errors := "Missing 'id' parameter"
		response := Response{
			Status: "Bad Request",
			Code:   400,
			Errors: errors,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	id, _ := strconv.Atoi(idStr)
	database.DB.Exec("DELETE FROM users WHERE id = ?", id)
	response := Response{
		Status: "OK",
		Code:   200,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
