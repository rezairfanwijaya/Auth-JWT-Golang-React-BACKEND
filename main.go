package main

import (
	"github.com/gofiber/fiber/v2"
	database "github.com/rezairfanwijaya/Auth-JWT-Golang-React/database"
	routes "github.com/rezairfanwijaya/Auth-JWT-Golang-React/routes"
)

func main() {
	// import koneksi database
	database.Connect()

	// kita menjalankan server menggunakan fiber biasanya menggunakan gin
	app := fiber.New()

	// import routes
	routes.Setup(app)
	

	// run server
	app.Listen(":8080")
}

