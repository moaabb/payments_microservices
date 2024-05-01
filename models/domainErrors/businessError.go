package domainErrors

import "net/http"

type BusinessError struct {
	Code       string
	Message    string
	StatusCode int
}

func NewBusinessError(code string, message string, status int) *BusinessError {
	return &BusinessError{
		Code:       code,
		Message:    message,
		StatusCode: status,
	}
}

// Errors
var (
	NotFoundError       = NewBusinessError("0001", "resource not found", http.StatusNotFound)
	InternalServerError = NewBusinessError("9999", "internal server error", http.StatusInternalServerError)
)
