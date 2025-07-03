package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `gorm:"uniqueIndex" json:"email"`
	Age   uint   `json:"age"`
}
