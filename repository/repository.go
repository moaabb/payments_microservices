package repository

type CustomerRepository interface {
	GetCustomers() []any
	CreateCustomer() any
	UpdateCustomer() any
}
