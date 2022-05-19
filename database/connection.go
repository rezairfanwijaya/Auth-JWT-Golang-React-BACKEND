package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	// koneksi ke databse
	dsn := "root:@tcp(localhost:3306)/jwt_go_react?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// cek error koneksi ke database
	if err != nil {
		panic(err)
	}
}
