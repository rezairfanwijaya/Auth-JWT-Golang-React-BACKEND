package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	database "github.com/rezairfanwijaya/Auth-JWT-Golang-React/database"
	routes "github.com/rezairfanwijaya/Auth-JWT-Golang-React/routes"
)

func main() {
	// import koneksi database
	database.Connect()

	// kita menjalankan server menggunakan fiber biasanya menggunakan gin
	app := fiber.New()

	// ini digunakan agar frontend bisa mengakses backend seperti mengambil cookie dan juga mengkases meskipun beda port (selain 8080 yg sudah di daftarkan dibawah)
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// import routes
	routes.Setup(app)

	// run server
	app.Listen(":8080")
}
