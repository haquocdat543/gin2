package main

import (
	"gin/src/config"
	"gin/src/db"
	"gin/src/router"
	"go.uber.org/zap"
)

func main() {

	// Logging
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Failed to initialize Zap logger: " + err.Error()) // Panic if initialization fails.
	}
	defer logger.Sync()

	config.LoadEnv()
	dbConn := config.InitDB()

	if err := db.Migrate(dbConn); err != nil {
		panic("Migration failed: " + err.Error())
	}

	r := router.SetupRouter(dbConn, logger)
	r.Run(":8080")
}
