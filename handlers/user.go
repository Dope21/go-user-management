package handlers

import (
	"encoding/json"
	"net/http"
	msg "user-management/constants/messages"
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
		utils.LogError(r, err.Error())
		utils.InvalidJSON(w)
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

	err = repository.InsertUser(user)
	if err != nil {
		utils.LogError(r, err.Error())
		utils.InternalServerError(w)
		return
	}

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

	var user models.UpdateUser
	user.ID = userID

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.LogError(r, err.Error())
		utils.InvalidJSON(w)
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