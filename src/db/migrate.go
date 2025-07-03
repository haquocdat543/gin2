package db

import (
	"gin/src/module/test"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&test.Test{},
		// Add more models here as your app grows
	)
}
