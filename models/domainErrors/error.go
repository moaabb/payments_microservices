package domainErrors

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	XValidator struct {
		validator *validator.Validate
		logger    *zap.Logger
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

func (xv XValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := xv.validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			xv.logger.Info(fmt.Sprintf("error validatind input: %s", elem.Stringify()))
			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func (m ErrorResponse) Stringify() string {
	out, _ := json.Marshal(m)

	return string(out)
}

func NewValidator(l *zap.Logger, v *validator.Validate) *XValidator {
	return &XValidator{
		logger:    l,
		validator: v,
	}
}
