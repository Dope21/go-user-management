package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	msg "user-management/constants/messages"
	"user-management/dto"
	"user-management/libs"
	"user-management/repository"
)

var ONE_WEEK_IN_MINS = 7 * 24 * 60 * 60
var TOKEN_PREFIX = "Bearer "

func Login(w http.ResponseWriter, r *http.Request) {

	var login dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.InvalidJSON(w)
		return
	}

	user, err := repository.GetUserByEmail(login.Email)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.ErrorResponse(w, http.StatusUnauthorized, msg.ErrLogin, nil)
		return
	}

	if user == nil {
		libs.LogError(r, msg.ErrNotFound)
		libs.ErrorResponse(w, http.StatusUnauthorized, msg.ErrLogin, nil)
		return
	}

	err = libs.ComparePassword(login.Password, user.Password)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.ErrorResponse(w, http.StatusUnauthorized, msg.ErrLogin, nil)
		return
	}

	accessToken, refreshToken, err := libs.GenerateTokens(user.ID)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.InternalServerError(w)
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

	token := dto.TokenResponse{
		AccessToken: accessToken,
	}

	libs.LogInfo(r, msg.SuccessLogin)
	libs.SuccessResponse(w, http.StatusOK, msg.SuccessLogin, token)
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
    if err != nil {
			libs.LogError(r, err.Error())
			switch {
			case errors.Is(err, http.ErrNoCookie):
				libs.InvalidToken(w)
			default:
				libs.InternalServerError(w)
			}
			return
    }
	
	userID, err := libs.VerifyToken(cookie.Value)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.InvalidToken(w)
		return
	}

	if _, err = repository.GetUserByID(userID); err != nil {
		libs.LogError(r, err.Error())
		libs.InternalServerError(w)
		return
	}

	accessToken, refreshToken, err := libs.GenerateTokens(userID)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.InternalServerError(w)
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

	token := dto.TokenResponse{
		AccessToken: accessToken,
	}

	libs.LogInfo(r, msg.SuccessLogin)
	libs.SuccessResponse(w, http.StatusOK, msg.SuccessLogin, token)
}

func VerifyToken(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, TOKEN_PREFIX)

	if authHeader == "" || token == authHeader {
		libs.LogError(r, msg.ErrNoToken)
		libs.NoToken(w)
		return
	}

	userID, err := libs.VerifyToken(token)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.InvalidToken(w)
		return
	}

	if _, err = repository.GetUserByID(userID); err != nil {
		libs.LogError(r, err.Error())
		libs.InternalServerError(w)
		return
	}

	libs.LogInfo(r, msg.SuccessGeneral)
	libs.SuccessResponse(w, http.StatusOK, msg.SuccessGeneral, nil)
}