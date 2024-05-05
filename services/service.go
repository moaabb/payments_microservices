package services

import (
	"github.com/jackc/pgx/v5"
	logging "github.com/moaabb/payments_microservices/customer/logger"
	"github.com/moaabb/payments_microservices/customer/models/domainErrors"
	"github.com/moaabb/payments_microservices/customer/models/entities"
	use_case "github.com/moaabb/payments_microservices/customer/usecase"
)

var logger = logging.GetLogger()

type CustomerService struct {
	db use_case.CustomerRepository
}

func NewCustomerService(db use_case.CustomerRepository) *CustomerService {
	return &CustomerService{
		db: db,
	}
}

func (m *CustomerService) GetCustomers() ([]entities.Customer, *domainErrors.BusinessError) {
	customers, err := m.db.GetCustomers()
	if err != nil {
		logger.Info(err.Error())
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
		logger.Info(err.Error())
		return nil, m.handleError(err)
	}
	return customer, nil
}

func (m *CustomerService) CreateCustomer(payload entities.Customer) (*entities.Customer, *domainErrors.BusinessError) {
	customer, err := m.db.CreateCustomer(payload)
	if err != nil {
		logger.Info(err.Error())
		return nil, m.handleError(err)
	}

	return customer, nil
}
func (m *CustomerService) UpdateCustomer(payload entities.Customer) (*entities.Customer, *domainErrors.BusinessError) {
	customer, err := m.db.UpdateCustomer(payload)
	if err != nil {
		logger.Info(err.Error())
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
