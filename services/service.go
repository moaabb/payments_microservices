package services

import (
	"github.com/moaabb/payments_microservices/customer/models/entities"
	"github.com/moaabb/payments_microservices/customer/repository"
	"go.uber.org/zap"
)

type CustomerService struct {
	db     repository.CustomerRepository
	logger *zap.Logger
}

func NewCustomerService(db repository.CustomerRepository, logger *zap.Logger) *CustomerService {
	return &CustomerService{
		db:     db,
		logger: logger,
	}
}

func (m *CustomerService) GetCustomers() ([]entities.Customer, error) {
	return m.db.GetCustomers()
}

func (m *CustomerService) GetCustomerById(id int) (*entities.Customer, error) {
	return m.db.GetCustomerById(id)
}

func (m *CustomerService) CreateCustomer(customer entities.Customer) (*entities.Customer, error) {
	return m.db.CreateCustomer(customer)
}
func (m *CustomerService) UpdateCustomer(customer entities.Customer) (*entities.Customer, error) {
	return nil, nil
}
