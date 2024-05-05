package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/moaabb/payments_microservices/customer/models/domainErrors"
	"github.com/moaabb/payments_microservices/customer/models/entities"
	use_case "github.com/moaabb/payments_microservices/customer/usecase"
	"go.opentelemetry.io/otel/trace"

	"go.uber.org/zap"
)

type CustomerHandler struct {
	service   use_case.CustomerService
	logger    *zap.Logger
	validator *domainErrors.XValidator
	tracer    trace.Tracer
}

func NewCustomerHandler(svc use_case.CustomerService, logger *zap.Logger, v *domainErrors.XValidator, t trace.Tracer) *CustomerHandler {
	return &CustomerHandler{
		service:   svc,
		logger:    logger,
		validator: v,
		tracer:    t,
	}
}

func (m *CustomerHandler) GetCustomers(c *fiber.Ctx) error {
	_, span := m.tracer.Start(c.Context(), fmt.Sprintf("%s %s", c.Method(), c.Request().URI().Path()))
	defer span.End()
	customers, err := m.service.GetCustomers()
	if err != nil {
		return c.Status(err.StatusCode).JSON(err)
	}

	return c.JSON(customers)
}

func (m *CustomerHandler) GetCustomerById(c *fiber.Ctx) error {
	_, span := m.tracer.Start(c.Context(), fmt.Sprintf("%s %s", c.Method(), c.Request().URI().Path()))
	defer span.End()
	customerId, _ := strconv.Atoi(c.Params("customerId"))

	customer, err := m.service.GetCustomerById(customerId)
	if err != nil {
		return c.Status(err.StatusCode).JSON(err)
	}

	return c.JSON(customer)
}

func (m *CustomerHandler) CreateCustomer(c *fiber.Ctx) error {
	_, span := m.tracer.Start(c.Context(), fmt.Sprintf("%s %s", c.Method(), c.Request().URI().Path()))
	defer span.End()
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
		return c.Status(err.StatusCode).JSON(err)
	}

	return c.JSON(customer)
}

func (m *CustomerHandler) UpdateCustomer(c *fiber.Ctx) error {
	return nil
}
