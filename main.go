package main

import (
	"fmt"
	"gin/src/db"
	"gin/src/router"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "develop", "effimatebackend", "postgres",
	)

	dbConn, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		panic("Failed to connect to database")
	}

	// Migrate schema
	if err := db.Migrate(dbConn); err != nil {
		panic("Failed to migrate database")
	}

	r := router.SetupRouter()
	r.Run(":8080")
}
