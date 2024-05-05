package services_test

import (
	"context"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/moaabb/payments_microservices/customer/models/domainErrors"
	"github.com/moaabb/payments_microservices/customer/models/entities"
	"github.com/moaabb/payments_microservices/customer/observability"
	"github.com/moaabb/payments_microservices/customer/services"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

type MockCustomerRepository struct {
	getCustomers    func(ctx context.Context) ([]entities.Customer, error)
	getCustomerById func(ctx context.Context, id int) (*entities.Customer, error)
	createCustomer  func(ctx context.Context, customer entities.Customer) (*entities.Customer, error)
	updateCustomer  func(ctx context.Context, customer entities.Customer) (*entities.Customer, error)
}

func (m *MockCustomerRepository) GetCustomers(ctx context.Context) ([]entities.Customer, error) {
	return m.getCustomers(ctx)
}
func (m *MockCustomerRepository) GetCustomerById(ctx context.Context, id int) (*entities.Customer, error) {
	return m.getCustomerById(ctx, id)
}

func (m *MockCustomerRepository) CreateCustomer(ctx context.Context, payload entities.Customer) (*entities.Customer, error) {
	return m.createCustomer(ctx, payload)
}
func (m *MockCustomerRepository) UpdateCustomer(ctx context.Context, payload entities.Customer) (*entities.Customer, error) {
	return m.updateCustomer(ctx, payload)
}

func TestGetCustomers(t *testing.T) {

}

func TestGetCustomerById(t *testing.T) {
	ctx := context.Background()

	// For testing to print out traces to the console
	// exp, err := newConsoleExporter()
	exp, err := observability.NewConsoleExporter()

	if err != nil {
		log.Fatal("failed to initialize exporter", zap.Error(err))
	}

	// Create a new tracer provider with a batch span processor and the given exporter.
	tp := observability.NewTraceProvider(exp)

	// Handle shutdown properly so nothing leaks.
	defer func() { _ = tp.Shutdown(ctx) }()

	otel.SetTracerProvider(tp)

	// Finally, set the tracer that can be used for this package.
	tracer := tp.Tracer("customer_svc")

	logger := zap.NewExample()
	date, _ := time.Parse("2006-02-1", "2006-02-1")
	expectedBody := entities.NewCustomer(1, "Teste", entities.Date{Time: date}, "teste@email.com", "77952658445")

	m := new(MockCustomerRepository)

	// Success case
	m = &MockCustomerRepository{
		getCustomerById: func(ctx context.Context, id int) (*entities.Customer, error) {
			return expectedBody, nil
		},
	}

	service := services.NewCustomerService(m, logger, tracer)
	resp, _ := service.GetCustomerById(ctx, *expectedBody.CustomerId)
	assert.Equal(t, expectedBody, resp, "they should be equal")

	// Not found case
	m = &MockCustomerRepository{
		getCustomerById: func(ctx context.Context, id int) (*entities.Customer, error) {
			return nil, pgx.ErrNoRows
		},
	}

	service = services.NewCustomerService(m, logger, tracer)
	_, httpErr := service.GetCustomerById(ctx, *expectedBody.CustomerId)

	assert.Equal(t, domainErrors.NotFoundError, err, "they should be equal")
	assert.Equal(t, http.StatusNotFound, httpErr.StatusCode, "they should be equal")
}
