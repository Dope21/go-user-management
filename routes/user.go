package routes

import (
	"net/http"
	"user-management/handlers"
)

func UserRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /user", handlers.GetAllUser)
	mux.HandleFunc("POST /user/register", handlers.RegisterUser)
}
