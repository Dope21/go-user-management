package routes

import (
	"encoding/json"
	"net/http"
)

func HealthCheckRouter(mux *http.ServeMux) {

	mux.HandleFunc("GET /readyz", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string { "status": "ready!" }
		json.NewEncoder(w).Encode(response)
	})

}