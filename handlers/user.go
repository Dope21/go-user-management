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

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	user, err := utils.ParsingBody[dto.CreateUserRequest](r)
	if err != nil {
		utils.LogError(r, err.Error())
		utils.InvalidJSON(w)
		return
	}

	fieldErrors, err := utils.ValidateBody(user)
	if err != nil {
		utils.LogError(r, err.Error())
		utils.InternalServerError(w)
		return
	}

	if fieldErrors != nil {
		utils.LogError(r, fmt.Sprint(fieldErrors))
		utils.InvalidBodyFields(w, fieldErrors)
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.LogError(r, err.Error())
		utils.InternalServerError(w)
		return
	}

	user.Password = hashedPassword

	result, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		utils.LogError(r, err.Error())
		utils.InternalServerError(w)
		return
	}

	if result != nil {
		utils.LogError(r, msg.ErrDuplicateEmail)
		utils.ErrorResponse(w, http.StatusBadRequest, msg.ErrDuplicateEmail, nil)
		return
	}

	newUser, err := repository.InsertUser(user)
	if err != nil {
		utils.LogError(r, err.Error())
		utils.InternalServerError(w)
		return
	}

	go func() {
    err := utils.SendEmailConfirmation(&newUser)
    if err != nil {
			utils.LogError(r, fmt.Sprintf(msg.ErrCantSendEmail, user.Email, err))
    }
	}()

	utils.LogInfo(r, msg.SuccessCreated)
	utils.SuccessResponse(w, http.StatusCreated, msg.SuccessCreated, nil)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	startRow := utils.ParseIntQueryParam(queryParams, "start")
	endRow := utils.ParseIntQueryParam(queryParams, "end")

	users, err := repository.GetAllUser(startRow, endRow)
	if err != nil {
		utils.LogError(r, err.Error())
		utils.InternalServerError(w)
		return
	}

	utils.LogInfo(r, msg.SuccessFetched)
	utils.SuccessResponse(w, http.StatusOK, msg.SuccessFetched, users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
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

	utils.LogInfo(r, msg.SuccessFetched)
	utils.SuccessResponse(w, http.StatusOK, msg.SuccessFetched, user)
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["user_id"]
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		utils.LogError(r, err.Error())
		utils.ErrorResponse(w, http.StatusBadRequest, msg.ErrInvalidUserID, nil)
		return
	}

	user, err := utils.ParsingBody[dto.UpdateUserRequest](r)
	if err != nil {
		utils.LogError(r, err.Error())
		utils.InvalidJSON(w)
		return
	}

	fieldError, err := utils.ValidateBody(user)
	if err != nil {
		utils.LogError(r, err.Error())
		utils.InternalServerError(w)
		return
	}

	if fieldError != nil {
		utils.LogError(r, fmt.Sprint(fieldError))
		utils.InvalidBodyFields(w, fieldError)
		return
	}

	if user.Password != nil {
		hashedPassword, err := utils.HashPassword(*user.Password)
		if err != nil {
			utils.LogError(r, err.Error())
			utils.InternalServerError(w)
			return
		}
		user.Password = &hashedPassword
	}

	user.ID = userID

	_, err = repository.UpdateUserByID(user)
	if err != nil {
		utils.LogError(r, err.Error())
		utils.InternalServerError(w)
		return
	}

	utils.LogInfo(r, msg.SuccessUpdated)
	utils.SuccessResponse(w, http.StatusOK, msg.SuccessUpdated, nil)
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["user_id"]
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		utils.LogError(r, err.Error())
		utils.ErrorResponse(w, http.StatusBadRequest, msg.ErrInvalidUserID, nil)
		return
	}

	err = repository.DeleteUserByID(userID)
	if err != nil {
		utils.LogError(r, err.Error())
		utils.InternalServerError(w)
		return
	}

	utils.LogInfo(r, msg.SuccessDeleted)
	utils.SuccessResponse(w, http.StatusOK, msg.SuccessDeleted, nil)
}