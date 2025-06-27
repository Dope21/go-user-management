package handlers

import (
	"fmt"
	"net/http"
	msg "user-management/constants/messages"
	"user-management/dto"
	"user-management/libs"
	"user-management/repository"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func ConfirmEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	userID, err :=libs.VerifyToken(token)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.InvalidToken(w)
		return
	}

	isActive := true
	user := dto.UpdateUserRequest{
		ID: userID,
		IsActive: &isActive,
	}

	_, err = repository.UpdateUserByID(&user)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.InternalServerError(w)
		return
	}

	libs.LogInfo(r, fmt.Sprintf(msg.SuccessConfirmEmail, userID))
	libs.SuccessResponse(w, http.StatusOK, msg.SuccessGeneral, nil)
}

func ResendCofirmEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["user_id"]

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.ErrorResponse(w, http.StatusBadRequest, msg.ErrInvalidUserID, nil)
		return
	}

	user, err := repository.GetUserByID(userID)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.InternalServerError(w)
		return
	}

	if user == nil {
		libs.LogError(r, msg.ErrNotFound)
		libs.NotFound(w)
		return
	}

	go func() {
    err := libs.SendEmailConfirmation(user)
    if err != nil {
			libs.LogError(r, fmt.Sprintf(msg.ErrCantSendEmail, user.Email, err))
    }
	}()

	libs.SuccessResponse(w, http.StatusOK, msg.SuccessSendEmail, nil)
}