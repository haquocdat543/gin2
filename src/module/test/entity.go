package test

import (
	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}
