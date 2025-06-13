package handlers

import (
	"fmt"
	"net/http"
	msg "user-management/constants/messages"
	"user-management/dto"
	"user-management/repository"
	"user-management/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func ConfirmEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	userID, err :=utils.VerifyToken(token)
	if err != nil {
		utils.LogError(r, err.Error())
		utils.InvalidToken(w)
		return
	}

	isActive := true
	user := dto.UpdateUserRequest{
		ID: userID,
		IsActive: &isActive,
	}

	_, err = repository.UpdateUserByID(&user)
	if err != nil {
		utils.LogError(r, err.Error())
		utils.InternalServerError(w)
		return
	}

	utils.LogInfo(r, fmt.Sprintf(msg.SuccessConfirmEmail, userID))
	utils.SuccessResponse(w, http.StatusOK, msg.SuccessGeneral, nil)
}

func ResendCofirmEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["user_id"]

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		utils.LogError(r, err.Error())
		utils.ErrorResponse(w, http.StatusBadRequest, msg.ErrInvalidUserID, nil)
		return
	}

	user, err := repository.GetUserByID(userID)
	if err != nil {
		utils.LogError(r, err.Error())
		utils.InternalServerError(w)
		return
	}

	if user == nil {
		utils.LogError(r, msg.ErrNotFound)
		utils.NotFound(w)
		return
	}

	go func() {
    err := utils.SendEmailConfirmation(user)
    if err != nil {
			utils.LogError(r, fmt.Sprintf(msg.ErrCantSendEmail, user.Email, err))
    }
	}()

	utils.SuccessResponse(w, http.StatusOK, msg.SuccessSendEmail, nil)
}