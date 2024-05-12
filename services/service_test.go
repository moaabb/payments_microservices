package services_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	logging "github.com/moaabb/payments_microservices/customer/logger"
	"github.com/moaabb/payments_microservices/customer/models/domainErrors"
	"github.com/moaabb/payments_microservices/customer/models/entities"
	"github.com/moaabb/payments_microservices/customer/services"
	"github.com/stretchr/testify/assert"
)

type MockCustomerRepository struct {
	getCustomers    func() ([]entities.Customer, error)
	getCustomerById func(id int) (*entities.Customer, error)
	createCustomer  func(customer entities.Customer) (*entities.Customer, error)
	updateCustomer  func(customer entities.Customer) (*entities.Customer, error)
}

func (m *MockCustomerRepository) GetCustomers() ([]entities.Customer, error) {
	return m.getCustomers()
}
func (m *MockCustomerRepository) GetCustomerById(id int) (*entities.Customer, error) {
	return m.getCustomerById(id)
}

func (m *MockCustomerRepository) CreateCustomer(payload entities.Customer) (*entities.Customer, error) {
	return m.createCustomer(payload)
}
func (m *MockCustomerRepository) UpdateCustomer(payload entities.Customer) (*entities.Customer, error) {
	return m.updateCustomer(payload)
}

func TestGetCustomerById(t *testing.T) {
	logging.InitLogger("INFO", "customer_svc", "test")

	date, _ := time.Parse("2006-02-1", "2006-02-1")
	expectedBody := entities.NewCustomer(1, "Teste", entities.Date{Time: date}, "teste@email.com", "77952658445")
	ctx := context.Background()

	m := new(MockCustomerRepository)

	// Success case
	m = &MockCustomerRepository{
		getCustomerById: func(id int) (*entities.Customer, error) {
			return expectedBody, nil
		},
	}

	service := services.NewCustomerService(m)
	resp, _ := service.GetCustomerById(ctx, *expectedBody.CustomerId)
	assert.Equal(t, expectedBody, resp, "they should be equal")

	// Not found case
	m = &MockCustomerRepository{
		getCustomerById: func(id int) (*entities.Customer, error) {
			return nil, pgx.ErrNoRows
		},
	}

	service = services.NewCustomerService(m)
	_, httpErr := service.GetCustomerById(ctx, *expectedBody.CustomerId)

	assert.Equal(t, domainErrors.NotFoundError, httpErr, "they should be equal")
	assert.Equal(t, http.StatusNotFound, httpErr.StatusCode, "they should be equal")
}
