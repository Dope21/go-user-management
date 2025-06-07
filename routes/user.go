package routes

import (
	"net/http"
	"user-management/handlers"

	"github.com/gorilla/mux"
)

func UserRouter(r *mux.Router) {
	userR := r.PathPrefix("/users").Subrouter()
	userR.HandleFunc("", handlers.GetAllUser).Methods(http.MethodGet)
	userR.HandleFunc("", handlers.RegisterUser).Methods(http.MethodPost)
	userR.HandleFunc("/{user_id}", handlers.UpdateUserByID).Methods(http.MethodPatch)
	userR.HandleFunc("/{user_id}", handlers.GetUserByID).Methods(http.MethodGet)
	userR.HandleFunc("/{user_id}", handlers.DeleteUserByID).Methods(http.MethodDelete)
	// userR.Use(utils.AuthenAdminMiddleware)
}
