package services

import (
	"github.com/jackc/pgx/v5"
	"github.com/moaabb/payments_microservices/customer/models/domainErrors"
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

func (m *CustomerService) GetCustomers() ([]entities.Customer, *domainErrors.BusinessError) {
	customers, err := m.db.GetCustomers()
	if err != nil {
		m.logger.Error(err.Error())
		return nil, m.handleError(err)
	}

	if customers == nil {
		customers = []entities.Customer{}
	}

	return customers, nil
}

func (m *CustomerService) GetCustomerById(id int) (*entities.Customer, *domainErrors.BusinessError) {
	customer, err := m.db.GetCustomerById(id)
	if err != nil {
		m.logger.Error(err.Error())
		return nil, m.handleError(err)
	}
	return customer, nil
}

func (m *CustomerService) CreateCustomer(payload entities.Customer) (*entities.Customer, *domainErrors.BusinessError) {
	customer, err := m.db.CreateCustomer(payload)
	if err != nil {
		m.logger.Error(err.Error())
		return nil, m.handleError(err)
	}

	return customer, nil
}
func (m *CustomerService) UpdateCustomer(payload entities.Customer) (*entities.Customer, *domainErrors.BusinessError) {
	customer, err := m.db.UpdateCustomer(payload)
	if err != nil {
		m.logger.Error(err.Error())
		return nil, m.handleError(err)
	}

	return customer, nil
}

func (m *CustomerService) handleError(err error) *domainErrors.BusinessError {
	switch err.Error() {
	case pgx.ErrNoRows.Error():
		return domainErrors.NotFoundError
	default:
		return domainErrors.InternalServerError
	}
}
