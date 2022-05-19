package database

import (
	model "github.com/rezairfanwijaya/Auth-JWT-Golang-React/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// bikin global varibale yang akan menampung koneksi ke database yg digunakan untuk proses crud di controller
var DB *gorm.DB

func Connect() {
	// koneksi ke databse
	dsn := "root:@tcp(localhost:3306)/jwt_go_react?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// cek error koneksi ke database
	if err != nil {
		panic(err)
	}

	// assign ke global variable
	DB = conn

	// jika tidak ada error maka jalankan migration tabel
	conn.AutoMigrate(&model.User{})
}
