package routes

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	HealthCheckRouter(r)
	UserRouter(r)

	return r
}
