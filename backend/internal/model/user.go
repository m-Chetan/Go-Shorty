package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"unique" validate:"required,min=3,max=50"`
	Email    string `gorm:"unique" validate:"required,email"`
	Password string `gorm:"required"`
}
