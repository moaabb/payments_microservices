package services

import (
	"context"

	"github.com/jackc/pgx/v5"
	logging "github.com/moaabb/payments_microservices/customer/logger"
	"github.com/moaabb/payments_microservices/customer/models/domainErrors"
	"github.com/moaabb/payments_microservices/customer/models/entities"
	use_case "github.com/moaabb/payments_microservices/customer/usecase"
)

type CustomerService struct {
	db     use_case.CustomerRepository
	logger *logging.ApplicationLogger
}

func NewCustomerService(db use_case.CustomerRepository) *CustomerService {
	return &CustomerService{
		db:     db,
		logger: logging.GetLogger(),
	}
}

func (m *CustomerService) GetCustomers(ctx context.Context) ([]entities.Customer, *domainErrors.BusinessError) {
	m.logger.WithContext(ctx).Info("retrieving customers from database")

	customers, err := m.db.GetCustomers()
	if err != nil {
		m.logger.Info(err.Error())
		return nil, m.handleError(err)
	}

	if customers == nil {
		customers = []entities.Customer{}
	}

	m.logger.WithContext(ctx).Infof("%d customers found", len(customers))

	return customers, nil
}

func (m *CustomerService) GetCustomerById(ctx context.Context, id int) (*entities.Customer, *domainErrors.BusinessError) {
	m.logger.WithContext(ctx).Infof("looking for customer id %d in database", id)
	customer, err := m.db.GetCustomerById(id)
	if err != nil {
		m.logger.Info(err.Error())
		return nil, m.handleError(err)
	}

	m.logger.WithContext(ctx).Infof("customer id %d found", id)
	return customer, nil
}

func (m *CustomerService) CreateCustomer(ctx context.Context, payload entities.Customer) (*entities.Customer, *domainErrors.BusinessError) {
	m.logger.WithContext(ctx).Infof("Creating customer record in database with the following payload, %s", payload.ToString())
	customer, err := m.db.CreateCustomer(payload)
	if err != nil {
		m.logger.Info(err.Error())
		return nil, m.handleError(err)
	}

	return customer, nil
}
func (m *CustomerService) UpdateCustomer(ctx context.Context, payload entities.Customer) (*entities.Customer, *domainErrors.BusinessError) {
	customer, err := m.db.UpdateCustomer(payload)
	if err != nil {
		m.logger.Info(err.Error())
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
