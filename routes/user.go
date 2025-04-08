package routes

import (
	"net/http"
	"user-management/handlers"

	"github.com/gorilla/mux"
)

func UserRouter(r *mux.Router) {
	r.HandleFunc("/user", handlers.GetAllUser).Methods(http.MethodGet)
	r.HandleFunc("/user/register", handlers.RegisterUser).Methods(http.MethodPost)
	r.HandleFunc("/user/{user_id}", handlers.UpdateUserByID).Methods(http.MethodPatch)
	r.HandleFunc("/user/{user_id}", handlers.GetUserByID).Methods(http.MethodGet)
}
