package router

import (
	"gin/src/module/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"go.uber.org/zap"
)

func SetupRouter(
	db *gorm.DB,
	logger *zap.Logger,
) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	testRepo := user.NewRepository(db)
	testService := user.NewService(testRepo)
	testHandler := user.NewHandler(testService)
	testHandler.RegisterRoutes(api, logger)

	return r
}
