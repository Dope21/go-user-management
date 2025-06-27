package libs

import (
	"encoding/json"
	"net/http"
	msg "user-management/constants/messages"
)

func ValidateRequestMethod(w http.ResponseWriter, r *http.Request, method string) {
	if r.Method != method {
		http.Error(w, msg.ErrInvalidMethod, http.StatusMethodNotAllowed)
	}
}

func ParsingBody[T any](r *http.Request) (*T, error) {
	var data T

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}