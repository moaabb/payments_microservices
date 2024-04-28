package services

import (
	"github.com/moaabb/payments_microservices/customer/repository"
	"go.uber.org/zap"
)

type CustomerService struct {
	db     repository.CustomerRepository
	logger *zap.Logger
}

func NewCustomerService(db repository.CustomerRepository, logger *zap.Logger) *CustomerService {

	return &CustomerService{
		db:     db,
		logger: logger,
	}
}
