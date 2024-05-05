package use_case

import (
	"github.com/moaabb/payments_microservices/customer/models/entities"
)

type CustomerRepository interface {
	GetCustomers() ([]entities.Customer, error)
	GetCustomerById(id int) (*entities.Customer, error)
	CreateCustomer(payload entities.Customer) (*entities.Customer, error)
	UpdateCustomer(payload entities.Customer) (*entities.Customer, error)
}
