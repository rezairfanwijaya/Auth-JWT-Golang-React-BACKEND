package controllers

import "github.com/gofiber/fiber/v2"

func AuthHandler(c *fiber.Ctx) error {
	return c.SendString("Hello Saya dari auth controller")

}
