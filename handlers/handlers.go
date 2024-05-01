package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/moaabb/payments_microservices/customer/models/domainErrors"
	"github.com/moaabb/payments_microservices/customer/models/entities"
	"github.com/moaabb/payments_microservices/customer/repository"
	"github.com/moaabb/payments_microservices/customer/services"
	"go.uber.org/zap"
)

type CustomerHandler struct {
	service   *services.CustomerService
	logger    *zap.Logger
	validator *domainErrors.XValidator
}

func NewCustomerHandler(db repository.CustomerRepository, logger *zap.Logger, v *domainErrors.XValidator) *CustomerHandler {
	service := services.NewCustomerService(db, logger)
	return &CustomerHandler{
		service:   service,
		logger:    logger,
		validator: v,
	}
}

func (m *CustomerHandler) GetCustomers(c *fiber.Ctx) error {
	customers, err := m.service.GetCustomers()
	if err != nil {
		m.logger.Error(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(customers)
}

func (m *CustomerHandler) GetCustomerById(c *fiber.Ctx) error {
	customerId, _ := strconv.Atoi(c.Params("customerId"))

	customer, err := m.service.GetCustomerById(customerId)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{})
	}

	return c.JSON(customer)
}

func (m *CustomerHandler) CreateCustomer(c *fiber.Ctx) error {
	var payload entities.Customer

	c.BodyParser(&payload)

	errs := m.validator.Validate(payload)
	if len(errs) > 0 {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		return &fiber.Error{
			Code:    fiber.ErrUnprocessableEntity.Code,
			Message: strings.Join(errMsgs, " and "),
		}

	}

	customer, err := m.service.CreateCustomer(payload)
	if err != nil {
		m.logger.Error(err.Error())
		c.Status(http.StatusUnprocessableEntity).JSON(err)
	}

	return c.JSON(customer)
}

func (m *CustomerHandler) UpdateCustomer(c *fiber.Ctx) error {
	return nil
}
