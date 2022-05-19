package models

// ini adalah struct representasi dari table user pada database
type User struct {
	Id       int    `json:"id" gorm:"primary-key"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"` // karena hasil encrypt password harus dalam bentuk byte
}
