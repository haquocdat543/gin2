package router

import (
	"gin/pkg/module/user"
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
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	userHandler.RegisterRoutes(api, logger)

	return r

}
