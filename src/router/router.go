package router

import (
	"gorm.io/gorm"
	"gin/src/module/user"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	testRepo := user.NewRepository(db)
	testService := user.NewService(testRepo)
	testHandler := user.NewHandler(testService)
	testHandler.RegisterRoutes(api)

	return r
}
