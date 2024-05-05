package domainErrors

import (
	"encoding/json"
	"net/http"
)

type BusinessError struct {
	Code       string `json:"errorCode"`
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

func NewBusinessError(code string, message string, status int) *BusinessError {
	return &BusinessError{
		Code:       code,
		Message:    message,
		StatusCode: status,
	}
}

func (m *BusinessError) ToString() string {
	out, _ := json.Marshal(m)

	return string(out)
}

// Errors
var (
	NotFoundError       = NewBusinessError("0001", "resource not found", http.StatusNotFound)
	InternalServerError = NewBusinessError("9999", "internal server error", http.StatusInternalServerError)
)
