package controllers

import (
	"github.com/gofiber/fiber/v2"
	database "github.com/rezairfanwijaya/Auth-JWT-Golang-React/database"
	model "github.com/rezairfanwijaya/Auth-JWT-Golang-React/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	// variable penampung dari data yang dikirim user dari method post
	var data map[string]string

	// parse map data tadi
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	// jika tidak terjadi error saat proses parse maka kita parsing inputan user ke model user yang sudah kita tulis di folder models

	// cek duplikasi email
	if database.DB.Where("email = ?", data["email"]).First(&model.User{}); database.DB.Error != nil {
		c.SendString("Email sudah terdaftar")
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Email already exist",
			"code":    400,
		})
	}

	// jika lolos cek duplikasi
	// namun khusus untuk password harus di encrypt terlebih dahulu, kita akan menggunakan package bcrypt dari golang
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := model.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	// save ke database
	database.DB.Create(&user)

	// maka return json dari data yang di input user
	c.SendStatus(200)
	return c.JSON(user)
}
