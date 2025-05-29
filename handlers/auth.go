package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	msg "user-management/constants/messages"
	"user-management/models"
	"user-management/repository"
	"user-management/utils"
)

var ONE_WEEK_IN_MINS = 7 * 24 * 60 * 60

func Login(w http.ResponseWriter, r *http.Request) {

	var login models.Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, msg.ErrMsgInvalidJSON, http.StatusBadRequest)
		return
	}

	user, err := repository.GetUserByEmail(login.Email)
	if err != nil {
		http.Error(w, msg.ErrMsgLogin, http.StatusUnauthorized)
		return
	}

	if !utils.ComparePassword(login.Password, user.Password) {
		http.Error(w, msg.ErrMsgLogin, http.StatusUnauthorized)
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(user.ID)
	if err != nil {
		http.Error(w, msg.ErrMsgInternalServer, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "refresh_token",
		Value: refreshToken,
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteNoneMode,
		Path: "/",
		MaxAge: ONE_WEEK_IN_MINS,
	})

	token := models.JWTTokens{
		AccessToken: accessToken,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
    if err != nil {
			switch {
			case errors.Is(err, http.ErrNoCookie):
				http.Error(w, msg.ErrMsgNoCookie, http.StatusBadRequest)
			default:
				http.Error(w, msg.ErrMsgInternalServer, http.StatusInternalServerError)
			}
			return
    }
	
	userID, err := utils.VerifyToken(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err = repository.GetUserByID(userID); err != nil {
		http.Error(w, msg.ErrMsgInternalServer, http.StatusInternalServerError)
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(userID)
	if err != nil {
		http.Error(w, msg.ErrMsgInternalServer, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "refresh_token",
		Value: refreshToken,
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteNoneMode,
		Path: "/",
		MaxAge: ONE_WEEK_IN_MINS,
	})

	token := models.JWTTokens{
		AccessToken: accessToken,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}