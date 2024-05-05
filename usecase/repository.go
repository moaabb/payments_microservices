package use_case

import (
	"context"

	"github.com/moaabb/payments_microservices/customer/models/entities"
)

type CustomerRepository interface {
	GetCustomers(ctx context.Context) ([]entities.Customer, error)
	GetCustomerById(ctx context.Context, id int) (*entities.Customer, error)
	CreateCustomer(ctx context.Context, payload entities.Customer) (*entities.Customer, error)
	UpdateCustomer(ctx context.Context, payload entities.Customer) (*entities.Customer, error)
}
