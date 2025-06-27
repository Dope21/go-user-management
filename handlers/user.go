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

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	user, err := libs.ParsingBody[dto.CreateUserRequest](r)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.InvalidJSON(w)
		return
	}

	fieldErrors, err := libs.ValidateBody(user)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.InternalServerError(w)
		return
	}

	if fieldErrors != nil {
		libs.LogError(r, fmt.Sprint(fieldErrors))
		libs.InvalidBodyFields(w, fieldErrors)
		return
	}

	hashedPassword, err := libs.HashPassword(user.Password)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.InternalServerError(w)
		return
	}

	user.Password = hashedPassword

	result, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.InternalServerError(w)
		return
	}

	if result != nil {
		libs.LogError(r, msg.ErrDuplicateEmail)
		libs.ErrorResponse(w, http.StatusBadRequest, msg.ErrDuplicateEmail, nil)
		return
	}

	newUser, err := repository.InsertUser(user)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.InternalServerError(w)
		return
	}

	go func() {
    err := libs.SendEmailConfirmation(&newUser)
    if err != nil {
			libs.LogError(r, fmt.Sprintf(msg.ErrCantSendEmail, user.Email, err))
    }
	}()

	libs.LogInfo(r, msg.SuccessCreated)
	libs.SuccessResponse(w, http.StatusCreated, msg.SuccessCreated, nil)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	startRow := libs.ParseIntQueryParam(queryParams, "start")
	endRow := libs.ParseIntQueryParam(queryParams, "end")

	users, err := repository.GetAllUser(startRow, endRow)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.InternalServerError(w)
		return
	}

	libs.LogInfo(r, msg.SuccessFetched)
	libs.SuccessResponse(w, http.StatusOK, msg.SuccessFetched, users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
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

	libs.LogInfo(r, msg.SuccessFetched)
	libs.SuccessResponse(w, http.StatusOK, msg.SuccessFetched, user)
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["user_id"]
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.ErrorResponse(w, http.StatusBadRequest, msg.ErrInvalidUserID, nil)
		return
	}

	user, err := libs.ParsingBody[dto.UpdateUserRequest](r)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.InvalidJSON(w)
		return
	}

	fieldError, err := libs.ValidateBody(user)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.InternalServerError(w)
		return
	}

	if fieldError != nil {
		libs.LogError(r, fmt.Sprint(fieldError))
		libs.InvalidBodyFields(w, fieldError)
		return
	}

	if user.Password != nil {
		hashedPassword, err := libs.HashPassword(*user.Password)
		if err != nil {
			libs.LogError(r, err.Error())
			libs.InternalServerError(w)
			return
		}
		user.Password = &hashedPassword
	}

	user.ID = userID

	_, err = repository.UpdateUserByID(user)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.InternalServerError(w)
		return
	}

	libs.LogInfo(r, msg.SuccessUpdated)
	libs.SuccessResponse(w, http.StatusOK, msg.SuccessUpdated, nil)
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["user_id"]
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.ErrorResponse(w, http.StatusBadRequest, msg.ErrInvalidUserID, nil)
		return
	}

	err = repository.DeleteUserByID(userID)
	if err != nil {
		libs.LogError(r, err.Error())
		libs.InternalServerError(w)
		return
	}

	libs.LogInfo(r, msg.SuccessDeleted)
	libs.SuccessResponse(w, http.StatusOK, msg.SuccessDeleted, nil)
}