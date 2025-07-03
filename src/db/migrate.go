package db

import (
	"gin/src/module/test"
	"gin/src/module/test2"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&test.Test{},
		&test2.Test2{},
		// Add more models here as your app grows
	)
}
