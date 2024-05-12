package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/moaabb/payments_microservices/customer/handlers"
	logging "github.com/moaabb/payments_microservices/customer/logger"
	"go.opentelemetry.io/otel/trace"
)

func getRoutes(handler *handlers.CustomerHandler, tp trace.TracerProvider) *fiber.App {
	app := fiber.New()

	app.Use(otelfiber.Middleware(otelfiber.WithTracerProvider(tp)))
	app.Use(traceMiddleware)

	app.Use(func(c *fiber.Ctx) error {
		ctx := c.Context()

		requestIdentifier := fmt.Sprintf("%s %s %s", c.Method(), c.BaseURL(), c.Request().URI().Path())
		logger.WithContext(ctx).Info(requestIdentifier)

		if slices.Contains([]string{"POST", "PATCH", "PUT"}, strings.ToUpper(c.Method())) {
			logger.WithContext(ctx).Infof("processing event body: %s", string(c.Body()))
		}

		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "server alive"})
	})

	app.Post("/v1/api/customers", handler.CreateCustomer)
	app.Get("/v1/api/customers", handler.GetCustomers)
	app.Get("/v1/api/customers/:customerId<int>", handler.GetCustomerById)
	app.Put("/v1/api/customers/:customerId", handler.UpdateCustomer)

	return app
}

func traceMiddleware(c *fiber.Ctx) error {
	cid := c.Get("x-cid", uuid.NewString())
	c.Locals(string(logging.CORRELATION_ID), cid)
	c.Set("x-cid", cid)

	return c.Next()
}
