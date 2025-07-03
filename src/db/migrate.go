package db

import (
	"gin/src/module/user"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&user.User{},
		// Add more models here as your app grows
	)
}
