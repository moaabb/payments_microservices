package services_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/moaabb/payments_microservices/customer/models/domainErrors"
	"github.com/moaabb/payments_microservices/customer/models/entities"
	"github.com/moaabb/payments_microservices/customer/services"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
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

func TestGetCustomers(t *testing.T) {

}

func TestGetCustomerById(t *testing.T) {
	logger := zap.NewExample()
	date, _ := time.Parse("2006-02-1", "2006-02-1")
	expectedBody := entities.NewCustomer(1, "Teste", entities.Date{Time: date}, "teste@email.com", "77952658445")

	m := new(MockCustomerRepository)

	// Success case
	m = &MockCustomerRepository{
		getCustomerById: func(id int) (*entities.Customer, error) {
			return expectedBody, nil
		},
	}

	service := services.NewCustomerService(m, logger)
	resp, _ := service.GetCustomerById(*expectedBody.CustomerId)
	assert.Equal(t, expectedBody, resp, "they should be equal")

	// Not found case
	m = &MockCustomerRepository{
		getCustomerById: func(id int) (*entities.Customer, error) {
			return nil, pgx.ErrNoRows
		},
	}

	service = services.NewCustomerService(m, logger)
	_, httpErr := service.GetCustomerById(*expectedBody.CustomerId)

	assert.Equal(t, domainErrors.NotFoundError, httpErr, "they should be equal")
	assert.Equal(t, http.StatusNotFound, httpErr.StatusCode, "they should be equal")
}
