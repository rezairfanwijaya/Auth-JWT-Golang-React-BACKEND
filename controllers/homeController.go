package controllers

import (
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	database "github.com/rezairfanwijaya/Auth-JWT-Golang-React/database"
)

func ShowAllUSer(c *fiber.Ctx) error {

	// * ambil cookie
	cookie := c.Cookies("jwt")

	// * ambil isi cookie
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Key), nil
	})

	// ! jika token tidak ditemukan
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Tidak ada token",
			"code":    fiber.StatusBadRequest,
		})
	}

	// ! jika token tidak valid
	if !token.Valid {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Token tidak valid",
			"code":    fiber.StatusBadRequest,
		})
	}

	// * tampilkan semua data user
	user := database.GetAllUser()

	// * hitung jumlah data user
	count := len(user)
	return c.JSON(fiber.Map{
		"message": "Berhasil menampilkan semua data user",
		"code":    200,
		"total":   count,
		"data":    user,
	})

}
