package routes

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	HealthCheckRouter(r)
	AuthRouter(r)
	UserRouter(r)
	EmailRouter(r)

	return r
}
