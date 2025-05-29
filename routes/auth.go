package routes

import (
	"net/http"
	"user-management/handlers"

	"github.com/gorilla/mux"
)

func AuthRouter(r *mux.Router) {
	r.HandleFunc("/auth/login", handlers.Login).Methods(http.MethodPost)
	r.HandleFunc("/auth/refresh", handlers.Refresh).Methods(http.MethodGet)
}