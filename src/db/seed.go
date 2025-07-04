package db

import (
	"gorm.io/gorm"
	"gin/src/module/user"
)

func Seed(db *gorm.DB) error {
	// return nil
	db.Create(user.Seed)
	return nil
}
