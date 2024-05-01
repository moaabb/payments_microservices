package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/moaabb/payments_microservices/customer/handlers"
)

// var logger = logger.GetLogger() /*  */

func getRoutes(handler *handlers.CustomerHandler) *fiber.App {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		log.Info(fmt.Sprintf("%s %s %s", c.Method(), c.BaseURL(), c.Request().URI().Path()))

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
