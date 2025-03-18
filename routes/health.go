package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func HealthCheckRouter(r *mux.Router) {

	r.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string { "status": "ready!" }
		json.NewEncoder(w).Encode(response)
	}).Methods(http.MethodGet)

}