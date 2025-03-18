package routes

import "net/http"

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	HealthCheckRouter(mux)
	UserRouter(mux)

	return mux
}
