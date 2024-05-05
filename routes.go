package main

import (
	"fmt"

	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/moaabb/payments_microservices/customer/handlers"
	"go.opentelemetry.io/otel/trace"
)

func getRoutes(handler *handlers.CustomerHandler, tp trace.TracerProvider) *fiber.App {
	app := fiber.New()

	app.Use(otelfiber.Middleware(otelfiber.WithTracerProvider(tp)))

	app.Use(func(c *fiber.Ctx) error {
		requestIdentifier := fmt.Sprintf("%s %s %s", c.Method(), c.BaseURL(), c.Request().URI().Path())
		logger.Info(requestIdentifier)
		logger.Info(fmt.Sprintf("processing event: {\"body\": %s}", string(c.Body())))

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
