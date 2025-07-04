package main

import (
	"gin/src/cli"
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

	// CLI
	cli.ExucuteCLI()

	// Router
	r := router.SetupRouter(
		dbConn,
		logger,
	)
	r.Run(":8080")

}
