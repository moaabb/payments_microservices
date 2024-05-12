package handlers_test

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/moaabb/payments_microservices/customer/handlers"
	"github.com/moaabb/payments_microservices/customer/models/domainErrors"
	"github.com/moaabb/payments_microservices/customer/models/entities"
	"github.com/stretchr/testify/assert"
)

var date, _ = time.Parse("2006-02-1", "2006-02-1")
var expectedBody = entities.NewCustomer(1, "Teste", entities.Date{Time: date}, "teste@email.com", "77952658445")

type MockCustomerService struct {
	getCustomers    func(ctx context.Context) ([]entities.Customer, *domainErrors.BusinessError)
	getCustomerById func(ctx context.Context, id int) (*entities.Customer, *domainErrors.BusinessError)
	createCustomer  func(ctx context.Context, customer entities.Customer) (*entities.Customer, *domainErrors.BusinessError)
	updateCustomer  func(ctx context.Context, customer entities.Customer) (*entities.Customer, *domainErrors.BusinessError)
}

func (m *MockCustomerService) GetCustomers(ctx context.Context) ([]entities.Customer, *domainErrors.BusinessError) {
	return m.getCustomers(ctx)
}
func (m *MockCustomerService) GetCustomerById(ctx context.Context, id int) (*entities.Customer, *domainErrors.BusinessError) {
	return m.getCustomerById(ctx, id)
}

func (m *MockCustomerService) CreateCustomer(ctx context.Context, payload entities.Customer) (*entities.Customer, *domainErrors.BusinessError) {
	return m.createCustomer(ctx, payload)
}
func (m *MockCustomerService) UpdateCustomer(ctx context.Context, payload entities.Customer) (*entities.Customer, *domainErrors.BusinessError) {
	return m.updateCustomer(ctx, payload)
}

func TestGetCustomerById(t *testing.T) {
	var output entities.Customer
	validator := domainErrors.NewValidator(validator.New())

	m := new(MockCustomerService)

	// Success case
	m = &MockCustomerService{
		getCustomerById: func(ctx context.Context, id int) (*entities.Customer, *domainErrors.BusinessError) {
			return expectedBody, nil
		},
	}

	// http.Request
	req := httptest.NewRequest("GET", "http://localhost:8080/v1/api/customers/1", nil)
	req.Header.Set("Content-Type", "application/json")

	app := fiber.New()

	handler := handlers.NewCustomerHandler(m, validator)
	app.Get("v1/api/customers/:customerId<int>", handler.GetCustomerById)
	resp, _ := app.Test(req)

	json.NewDecoder(resp.Body).Decode(&output)
	assert.Equal(t, expectedBody.ToString(), output.ToString(), "they should be equal")

	// // Not found case
	m = &MockCustomerService{
		getCustomerById: func(ctx context.Context, id int) (*entities.Customer, *domainErrors.BusinessError) {
			return nil, domainErrors.NotFoundError
		},
	}

	// http.Request
	req = httptest.NewRequest("GET", "http://localhost:8080/v1/api/customers/1", nil)
	req.Header.Set("Content-Type", "application/json")

	app = fiber.New()

	var errorOut domainErrors.BusinessError

	handler = handlers.NewCustomerHandler(m, validator)
	app.Get("v1/api/customers/:customerId<int>", handler.GetCustomerById)
	resp, _ = app.Test(req)

	json.NewDecoder(resp.Body).Decode(&errorOut)
	assert.Equal(t, domainErrors.NotFoundError.ToString(), errorOut.ToString(), "they should be equal")

}
