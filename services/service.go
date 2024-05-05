package services

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/moaabb/payments_microservices/customer/models/domainErrors"
	"github.com/moaabb/payments_microservices/customer/models/entities"
	use_case "github.com/moaabb/payments_microservices/customer/usecase"
	"go.opentelemetry.io/otel/trace"

	"go.uber.org/zap"
)

type CustomerService struct {
	db     use_case.CustomerRepository
	logger *zap.Logger
	tracer trace.Tracer
}

func NewCustomerService(db use_case.CustomerRepository, logger *zap.Logger, t trace.Tracer) *CustomerService {
	return &CustomerService{
		db:     db,
		logger: logger,
		tracer: t,
	}
}

func (m *CustomerService) GetCustomers(ctx context.Context) ([]entities.Customer, *domainErrors.BusinessError) {
	ctx, span := m.tracer.Start(ctx, "CustomerService")
	defer span.End()
	customers, err := m.db.GetCustomers(ctx)
	if err != nil {
		m.logger.Info(err.Error())
		return nil, m.handleError(err)
	}

	if customers == nil {
		customers = []entities.Customer{}
	}

	return customers, nil
}

func (m *CustomerService) GetCustomerById(ctx context.Context, id int) (*entities.Customer, *domainErrors.BusinessError) {
	ctx, span := m.tracer.Start(ctx, "CustomerService")
	defer span.End()
	customer, err := m.db.GetCustomerById(ctx, id)
	if err != nil {
		m.logger.Info(err.Error())
		return nil, m.handleError(err)
	}
	return customer, nil
}

func (m *CustomerService) CreateCustomer(ctx context.Context, payload entities.Customer) (*entities.Customer, *domainErrors.BusinessError) {
	ctx, span := m.tracer.Start(ctx, "CustomerService")
	defer span.End()
	customer, err := m.db.CreateCustomer(ctx, payload)
	if err != nil {
		m.logger.Info(err.Error())
		return nil, m.handleError(err)
	}

	return customer, nil
}
func (m *CustomerService) UpdateCustomer(ctx context.Context, payload entities.Customer) (*entities.Customer, *domainErrors.BusinessError) {
	ctx, span := m.tracer.Start(ctx, "CustomerService")
	defer span.End()
	customer, err := m.db.UpdateCustomer(ctx, payload)
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
