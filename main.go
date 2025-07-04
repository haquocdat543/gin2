package main

import (
	"gin/src/config"
	"gin/src/db"
	"gin/src/router"
)

func main() {

	// Env
	config.InitEnv()

	// Logging
	logger := config.InitLog()

	// Database
	dbConn := config.InitDB()

	// Migration
	if err := db.Migrate(dbConn); err != nil {
		panic("Migration failed: " + err.Error())
	}

	// Seeding
	if err := db.Seed(dbConn); err != nil {
		panic("Seeding failed: " + err.Error())
	}

	// Router
	r := router.SetupRouter(
		dbConn,
		logger,
	)
	r.Run(":8080")

}
