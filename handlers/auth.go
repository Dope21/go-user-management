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
		utils.LogError(r, err)
		utils.InvalidJSON(w)
		return
	}

	user, err := repository.GetUserByEmail(login.Email)
	if err != nil {
		utils.LogError(r, err)
		utils.ErrorResponse(w, http.StatusUnauthorized, msg.ErrLogin, nil)
		return
	}

	err = utils.ComparePassword(login.Password, user.Password)
	if err != nil {
		utils.LogError(r, err)
		utils.ErrorResponse(w, http.StatusUnauthorized, msg.ErrLogin, nil)
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(user.ID)
	if err != nil {
		utils.LogError(r, err)
		utils.InternalServerError(w)
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

	utils.LogInfo(r, msg.SuccessLogin)
	utils.SuccessResponse(w, http.StatusOK, msg.SuccessLogin, token)
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
    if err != nil {
			utils.LogError(r, err)
			switch {
			case errors.Is(err, http.ErrNoCookie):
				utils.InvalidToken(w)
			default:
				utils.InternalServerError(w)
			}
			return
    }
	
	userID, err := utils.VerifyToken(cookie.Value)
	if err != nil {
		utils.LogError(r, err)
		utils.InvalidToken(w)
		return
	}

	if _, err = repository.GetUserByID(userID); err != nil {
		utils.LogError(r, err)
		utils.InternalServerError(w)
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(userID)
	if err != nil {
		utils.LogError(r, err)
		utils.InternalServerError(w)
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

	utils.LogInfo(r, msg.SuccessLogin)
	utils.SuccessResponse(w, http.StatusOK, msg.SuccessLogin, token)
}