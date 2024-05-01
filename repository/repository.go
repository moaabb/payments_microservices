package repository

import (
	"github.com/moaabb/payments_microservices/customer/models/entities"
)

type CustomerRepository interface {
	GetCustomers() ([]entities.Customer, error)
	GetCustomerById(id int) (*entities.Customer, error)
	CreateCustomer(customer entities.Customer) (*entities.Customer, error)
	UpdateCustomer(customer entities.Customer) (*entities.Customer, error)
}
