package routes

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/rezairfanwijaya/Auth-JWT-Golang-React/controllers"
)

func Setup(app *fiber.App) {
	// route
	app.Post("/api/register", controller.Register)
}
