package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"user-management/models"
	"user-management/repository"
	"user-management/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "sorry, something went wrong", http.StatusInternalServerError)
		return
	}

	user.Password = hashedPassword

	err = repository.InsertUser(user)
	if err != nil {
		http.Error(w, "sorry, something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "user registered successfully")
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	startRow := utils.ParseIntQueryParam(queryParams, "start")
	endRow := utils.ParseIntQueryParam(queryParams, "end")

	users, err := repository.GetAllUser(startRow, endRow)
	if err != nil {
		http.Error(w, "sorry, something went wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["user_id"]
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	user, err := repository.GetUserByID(userID)
	if err != nil {
		http.Error(w, "sorry, something went wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["user_id"]
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var user models.UpdateUser
	user.ID = userID

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	if user.Password != nil {
		hashedPassword, err := utils.HashPassword(*user.Password)
		if err != nil {
			http.Error(w, "sorry, something went wrong", http.StatusInternalServerError)
			return
		}
		user.Password = &hashedPassword
	}

	_, err = repository.UpdateUserByID(user)
	if err != nil {
		http.Error(w, "sorry, something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "user update successfully")
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["user_id"]
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	err = repository.DeleteUserByID(userID)
	if err != nil {
		log.Printf("error %v", err)
		http.Error(w, "sorry, something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "user delete successfully")
}