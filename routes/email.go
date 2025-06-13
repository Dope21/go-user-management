package routes

import (
	"net/http"
	"user-management/handlers"
	"user-management/utils"

	"github.com/gorilla/mux"
)

func EmailRouter(r *mux.Router) {
	emailR := r.PathPrefix("/email").Subrouter()
	emailR.HandleFunc("/confirm/{token}", handlers.ConfirmEmail).Methods(http.MethodGet)
	emailR.HandleFunc("/resend/{user_id}", handlers.ResendCofirmEmail).Methods(http.MethodGet)
	emailR.Use(utils.AuthenMiddleware)
}