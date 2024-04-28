package handlers

import (
	"github.com/moaabb/payments_microservices/customer/repository"
	"github.com/moaabb/payments_microservices/customer/services"
	"go.uber.org/zap"
)

type CustomerHandler struct {
	service *services.CustomerService
	logger  *zap.Logger
}

func NewCustomerHandler(db repository.CustomerRepository, logger *zap.Logger) *CustomerHandler {
	service := services.NewCustomerService(db, logger)
	return &CustomerHandler{
		service: service,
		logger:  logger,
	}
}
