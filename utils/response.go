package utils

import (
	"encoding/json"
	"net/http"
	msg "user-management/constants/messages"
	"user-management/models"
)

func InvalidJSON(w http.ResponseWriter) {
	statusCode := http.StatusBadRequest
	res := models.NewHTTPResponse(false, statusCode, msg.ErrInvalidJSON, nil)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}

func InvalidToken(w http.ResponseWriter) {
	statusCode := http.StatusUnauthorized
	res := models.NewHTTPResponse(false, statusCode, msg.ErrInvalidToken, nil)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}

func NoToken(w http.ResponseWriter) {
	statusCode := http.StatusUnauthorized
	res := models.NewHTTPResponse(false, statusCode, msg.ErrNoToken, nil)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}

func Forbidden(w http.ResponseWriter) {
	statusCode := http.StatusForbidden
	res := models.NewHTTPResponse(false, statusCode, msg.ErrForbidden, nil)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}

func InternalServerError(w http.ResponseWriter) {
	statusCode := http.StatusInternalServerError
	res := models.NewHTTPResponse(false, statusCode, msg.ErrInternalServer, nil)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}

func SuccessResponse(w http.ResponseWriter, statusCode int, message string, data any) {
	res := models.NewHTTPResponse(true, statusCode, message, data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}

func ErrorResponse(w http.ResponseWriter, statusCode int, message string, data any) {
	res := models.NewHTTPResponse(false, statusCode, message, data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}