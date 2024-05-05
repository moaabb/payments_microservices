package use_case

import (
	"context"

	"github.com/moaabb/payments_microservices/customer/models/domainErrors"
	"github.com/moaabb/payments_microservices/customer/models/entities"
)

type CustomerService interface {
	GetCustomers(ctx context.Context) ([]entities.Customer, *domainErrors.BusinessError)
	GetCustomerById(ctx context.Context, id int) (*entities.Customer, *domainErrors.BusinessError)
	CreateCustomer(ctx context.Context, payload entities.Customer) (*entities.Customer, *domainErrors.BusinessError)
	UpdateCustomer(ctx context.Context, payload entities.Customer) (*entities.Customer, *domainErrors.BusinessError)
}
