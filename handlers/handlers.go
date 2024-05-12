package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	logging "github.com/moaabb/payments_microservices/customer/logger"
	"github.com/moaabb/payments_microservices/customer/models/domainErrors"
	"github.com/moaabb/payments_microservices/customer/models/entities"
	use_case "github.com/moaabb/payments_microservices/customer/usecase"
)

type CustomerHandler struct {
	service   use_case.CustomerService
	validator *domainErrors.XValidator
	logger    *logging.ApplicationLogger
}

func NewCustomerHandler(svc use_case.CustomerService, v *domainErrors.XValidator) *CustomerHandler {
	return &CustomerHandler{
		service:   svc,
		validator: v,
		logger:    logging.GetLogger(),
	}
}

func (m *CustomerHandler) GetCustomers(c *fiber.Ctx) error {
	customers, err := m.service.GetCustomers(c.Context())
	if err != nil {
		return c.Status(err.StatusCode).JSON(err)
	}

	return c.JSON(customers)
}

func (m *CustomerHandler) GetCustomerById(c *fiber.Ctx) error {
	customerId, _ := strconv.Atoi(c.Params("customerId"))

	customer, err := m.service.GetCustomerById(c.Context(), customerId)
	if err != nil {
		return c.Status(err.StatusCode).JSON(err)
	}

	return c.JSON(customer)
}

func (m *CustomerHandler) CreateCustomer(c *fiber.Ctx) error {
	var payload entities.Customer

	c.BodyParser(&payload)
	errs := m.validator.Validate(c.Context(), payload)
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

	customer, err := m.service.CreateCustomer(c.Context(), payload)
	if err != nil {
		return c.Status(err.StatusCode).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(customer)
}

func (m *CustomerHandler) UpdateCustomer(c *fiber.Ctx) error {
	return nil
}
