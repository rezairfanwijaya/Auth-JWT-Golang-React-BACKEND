package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", Hello)

	app.Listen(":8080")
}

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
