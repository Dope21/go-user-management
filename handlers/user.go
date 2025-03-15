package handlers

import (
	"encoding/json"
	"net/http"
	"user-management/models"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	ecoder := json.NewEncoder(w)
	ecoder.Encode(&user)
}