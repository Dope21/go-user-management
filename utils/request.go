package utils

import (
	"net/http"
	msg "user-management/constants/messages"
)

func ValidateRequestMethod(w http.ResponseWriter, r *http.Request, method string) {
	if r.Method != method {
		http.Error(w, msg.ErrInvalidMethod, http.StatusMethodNotAllowed)
	}
}