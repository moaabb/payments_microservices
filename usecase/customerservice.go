package use_case

import (
	"github.com/moaabb/payments_microservices/customer/models/domainErrors"
	"github.com/moaabb/payments_microservices/customer/models/entities"
)

type CustomerService interface {
	GetCustomers() ([]entities.Customer, *domainErrors.BusinessError)
	GetCustomerById(id int) (*entities.Customer, *domainErrors.BusinessError)
	CreateCustomer(payload entities.Customer) (*entities.Customer, *domainErrors.BusinessError)
	UpdateCustomer(payload entities.Customer) (*entities.Customer, *domainErrors.BusinessError)
}
