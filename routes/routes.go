package routes

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/rezairfanwijaya/Auth-JWT-Golang-React/controllers"
)

func Setup(app *fiber.App) {
	// route
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Get("/api/user", controller.User)
	app.Get("/api/logout", controller.Logout)
	app.Get("/", controller.ShowAllUSer)
}
