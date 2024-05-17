package domainErrors

import (
	"context"
	"encoding/json"

	"github.com/go-playground/validator/v10"
	logging "github.com/moaabb/payments_microservices/customer/logger"
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
		logger    *logging.ApplicationLogger
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

func (xv XValidator) Validate(ctx context.Context, data interface{}) []ErrorResponse {
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

			xv.logger.WithContext(ctx).Infof("error validatind input: %s", elem.Stringify())
			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func (m ErrorResponse) Stringify() string {
	out, _ := json.Marshal(m)

	return string(out)
}

func NewValidator(v *validator.Validate) *XValidator {
	return &XValidator{
		validator: v,
		logger:    logging.GetLogger(),
	}
}
