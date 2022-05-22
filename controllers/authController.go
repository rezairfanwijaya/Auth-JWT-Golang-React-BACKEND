package controllers

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	database "github.com/rezairfanwijaya/Auth-JWT-Golang-React/database"
	model "github.com/rezairfanwijaya/Auth-JWT-Golang-React/models"
	"golang.org/x/crypto/bcrypt"
)

var Key = "mysecret123890"

// function handler register
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
	var user model.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id != 0 {
		c.Status(fiber.StatusConflict)
		return c.JSON(fiber.Map{
			"status":  "Conflict",
			"code":    409,
			"message": "Email sudah terdaftar",
		})
	}

	// jika lolos cek duplikasi
	// namun khusus untuk password harus di encrypt terlebih dahulu, kita akan menggunakan package bcrypt dari golang
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	newUser := model.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	// save ke database
	database.DB.Create(&newUser)

	// maka return json dari data yang di input user
	return c.JSON(fiber.Map{
		"message": "Registrasi berhasil",
		"code":    200,
		"data":    newUser,
	})
}

// function handler login
func Login(c *fiber.Ctx) error {
	// bikin variable penampung dari data yang dikirim user dari method post
	var data map[string]string

	// kita parse map data tadi
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	// kita insiasi models user
	var user model.User

	// cek apakah ada data di database dengan email yang diinputkan oleh user
	database.DB.Where("email = ?", data["email"]).First(&user)
	// jika tidak ada data yang ditemukan
	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  "Not Found",
			"code":    400,
			"message": "Pengguna tidak ditemukan, silahakan register terlebih dahulu di alamat http://localhost:3000/api/register",
		})
	}

	// jika ada data yang ditemukan
	// kita cek apakah password yang diinputkan oleh user sama dengan password yang ada di database
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Password salah",
			"code":    400,
		})
	}

	// ketika user berhasil login, buatkan jwt token
	// bikin claim, kita persiapkan dulu data yang akan kita masukan ke dalam claim(payload jwt)
	issuerID := strconv.Itoa(user.Id)
	exp := time.Now().Add(time.Hour * 24)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: issuerID,
		ExpiresAt: &jwt.Time{
			Time: exp,
		},
	})

	// lalu tandatangani token yang dibuat menggunakan secret Key yang sudah kita buat
	token, err := claims.SignedString([]byte(Key))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Token gagal dibuat",
			"code":    500,
		})
	}

	// ketika berhasil ditandatangai, kita masukan token jwt ke dalam cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  exp,
		HTTPOnly: true, // httponly digunakan agar nanti frontend tidak bisa mengakses(ubah) token, hanya bisa memilikinya saja
	}

	// lalu simpan cookie
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Berhasil login",
		"name":    user.Name,
	})
}

// function untuk mengecek user siapa yang sedang login
func User(c *fiber.Ctx) error {
	// kita ambil cookie
	cookie := c.Cookies("jwt") // kenapa jwt? karena kita set name cookie nya jwt, bisa cek difunction login

	// kita cek apakah jwt yang ada di cookie itu valid atau tidak
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		// jika valid
		return []byte(Key), nil
	})

	// cek error
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Token tidak valid",
			"code":    401,
		})
	}

	// ambil informasi yg ada didalam token untuk mengetahui siapa yang sedang login
	// karena di dalam token tadi kita menyisipkan id user
	claims := token.Claims.(*jwt.StandardClaims)
	var user model.User

	database.DB.Where("id = ?", claims.Issuer).First(&user) // issuer berisi id user sesuai dengan apa yg kita set di function login

	return c.JSON(user)

}

// function untuk logout
func Logout(c *fiber.Ctx) error {
	// untuk logout kita hanya perlu mengosongkan value cookie jwt dan memundurkan waktu menjadi minus
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(time.Hour * -1),
		HTTPOnly: true,
	}

	// simpan cookie
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Berhasil logout",
		"code":    fiber.StatusOK,
	})
}
