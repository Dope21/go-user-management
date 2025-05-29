package utils

import "net/http"

func ValidateRequestMethod(w http.ResponseWriter, r *http.Request, method string) {
	if r.Method != method {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
	}
}