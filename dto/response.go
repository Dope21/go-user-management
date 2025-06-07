package dto

type HTTPResponse struct {
	IsSuccess  bool   `json:"is_success"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

func NewHTTPResponse(isSuccess bool, statusCode int, message string, data any) HTTPResponse {
	if data == nil {
		data = map[string]any{}
	}

	return HTTPResponse{
		IsSuccess:  isSuccess,
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}