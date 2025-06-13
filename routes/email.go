package routes

import (
	"net/http"
	"user-management/handlers"

	"github.com/gorilla/mux"
)

func EmailRouter(r *mux.Router) {
	emailR := r.PathPrefix("/email").Subrouter()
	emailR.HandleFunc("/confirm/{token}", handlers.ConfirmEmail).Methods(http.MethodGet)
}