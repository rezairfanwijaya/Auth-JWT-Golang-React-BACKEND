package models

// ini adalah struct representasi dari table user pada database
type User struct {
	Id       int `gorm:"primary-key"`
	Name     string
	Email    string `gorm:"unique"`
	Password []byte // karena hasil encrypt password harus dalam bentuk byte
}
