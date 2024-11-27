package handlers

import (
	"encoding/json"
	"net/http"
	"rest-api-go/database"
	"rest-api-go/models"
	"strconv"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
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
		response := models.Response{
			Status: "Method Not Allowed",
			Code:   405,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
	}
}

func handleReadUsers(w http.ResponseWriter) {
	rows, _ := database.DB.Query("SELECT id, name FROM users")
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}

	response := models.Response{
		Status: "OK",
		Code:   200,
		Data:   users,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" && len(r.Form["name"]) == 0 {
		errors := []string{"Missing 'name' parameter"}
		response := models.Response{
			Status: "Bad Request",
			Code:   400,
			Errors: errors,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	newUser := models.User{Name: name}
	database.DB.Exec("INSERT INTO users (name) VALUES (?)", newUser.Name)

	response := models.Response{
		Status: "Created",
		Code:   201,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	name := r.FormValue("name")
	if idStr == "" && len(r.Form["id"]) == 0 || name == "" && len(r.Form["name"]) == 0 {
		errors := []string{"Missing 'id' or 'name' parameter"}
		response := models.Response{
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
	newUser := models.User{ID: id, Name: name}
	database.DB.Exec("UPDATE users SET name = ? WHERE id = ?", newUser.Name, newUser.ID)

	response := models.Response{
		Status: "OK",
		Code:   200,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" && len(r.Form["id"]) == 0 {
		errors := []string{"Missing 'id' parameter"}
		response := models.Response{
			Status: "Bad Request",
			Code:   400,
			Errors: errors,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	database.DB.Exec("DELETE FROM users WHERE id = ?", id)

	response := models.Response{
		Status: "OK",
		Code:   200,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
