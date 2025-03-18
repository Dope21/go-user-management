package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user-management/models"
	"user-management/repository"
	"user-management/utils"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	utils.ValidateRequestMethod(w, r, http.MethodPost)

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user.Password = hashedPassword

	err = repository.InsertUser(user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User registered successfully")
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	utils.ValidateRequestMethod(w, r, http.MethodGet)

	queryParams := r.URL.Query()
	startRow := utils.ParseIntQueryParam(queryParams, "start")
	endRow := utils.ParseIntQueryParam(queryParams, "end")

	users, err := repository.GetAllUser(startRow, endRow)
	if err != nil {
		http.Error(w, "Error query users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}