package routes

import (
	"net/http"
	"user-management/handlers"

	"github.com/gorilla/mux"
)

func UserRouter(r *mux.Router) {
	r.HandleFunc("/users", handlers.GetAllUser).Methods(http.MethodGet)
	r.HandleFunc("/users/register", handlers.RegisterUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{user_id}", handlers.UpdateUserByID).Methods(http.MethodPatch)
	r.HandleFunc("/users/{user_id}", handlers.GetUserByID).Methods(http.MethodGet)
	r.HandleFunc("/users/{user_id}", handlers.DeleteUserByID).Methods(http.MethodDelete)
}
