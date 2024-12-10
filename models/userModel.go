package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Phone    string `gorm:"unique"`
	Password string
	Username string
}