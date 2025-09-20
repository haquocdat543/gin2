package main

import (
	"gin/pkg/cli"
	"gin/pkg/config"
	"gin/pkg/db"
	"gin/pkg/router"
	"log"
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

	ginRunError := r.Run(":8080")
	if ginRunError != nil {
		log.Fatal(ginRunError)
	}

}
