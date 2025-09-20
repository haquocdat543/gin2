package config

import (
	"log"

	"go.uber.org/zap"
)

func InitLog() *zap.Logger {

	logger, err := zap.NewProduction()
	if err != nil {
		panic(
			"Failed to initialize Zap logger: " + err.Error(),
		) // Panic if initialization fails.
	}

	defer func() {
		logSyncError := logger.Sync()
		if logSyncError != nil {
			log.Print(logSyncError.Error())
		}
	}()

	return logger
}
