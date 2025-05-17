package controllers

import (
	"encoding/json"
	"go-mvc-app/models"
	"go-mvc-app/views"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	views.JSONResponse(w, users)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	models.AddUser(user)
	views.JSONResponse(w, map[string]string{"message": "User added successfully"})
}
