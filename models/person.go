package models

import "gorm.io/gorm"

type Person struct {
    gorm.Model
	Name     string `form:"name" validate:"required,min=2"`
	Email    string `form:"email" gorm:"unique" validate:"required,email"`
	Cell     string `form:"cell" validate:"required,min=9"`
	Password string `form:"password" validate:"required,min=8,max=12"`
	Role     string `form:"role"`
}
