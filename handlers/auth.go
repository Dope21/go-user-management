package handlers

import (
	"encoding/json"
	"net/http"
	"user-management/models"
	"user-management/repository"
	"user-management/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var login models.Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	user, err := repository.GetUserByEmail(login.Email)
	if err != nil {
		http.Error(w, "Error login user. Wrong email or password", http.StatusUnauthorized)
		return
	}

	if !utils.ComparePassword(login.Password, user.Password) {
		http.Error(w, "Error login user. Wrong email or password", http.StatusUnauthorized)
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(user.ID)
	if err != nil {
		http.Error(w, "Sorry, something went wrong.", http.StatusInternalServerError)
		return
	}

	tokens := models.JWTTokens{
		AccessToken: &accessToken,
		RefreshToken: &refreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokens)
}