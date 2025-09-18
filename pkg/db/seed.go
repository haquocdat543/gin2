package db

import (
	"gin/pkg/module/user"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	// return nil
	db.Create(user.Seed)
	return nil
}
