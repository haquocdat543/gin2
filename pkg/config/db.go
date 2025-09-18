package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/schema"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		ENV.DBHost,
		ENV.DBPort,
		ENV.DBUser,
		ENV.DBPassword,
		ENV.DBName,
	)

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	return db
}
