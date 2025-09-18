package config

import (
	"go.uber.org/zap"
)

func InitLog() *zap.Logger {

	logger, err := zap.NewProduction()
	if err != nil {
		panic(
			"Failed to initialize Zap logger: " + err.Error(),
		) // Panic if initialization fails.
	}
	defer logger.Sync()

	return logger
}
