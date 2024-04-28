package main

import (
	"github.com/gofiber/fiber/v2"
)

func getRoutes() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "server alive"})
	})

	return app
}
