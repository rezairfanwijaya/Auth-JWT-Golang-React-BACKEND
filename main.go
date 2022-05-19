package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// koneksi ke databse
	dsn := "root:@tcp(localhost:3306)/jwt_go_react?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// cek error koneksi ke database
	if err != nil {
		panic(err)
	}

	// kita menjalankan server menggunakan fiber biasanya menggunakan gin
	app := fiber.New()

	// route
	app.Get("/", Hello)

	// run server
	app.Listen(":8080")
}

// function handler
func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
