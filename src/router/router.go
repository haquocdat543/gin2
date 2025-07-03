package router

import (
	"gorm.io/gorm"
	"gin/src/module/user"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	// ✅ Initialize service and handler properly
	testRepo := user.NewRepository(db)
	testService := user.NewService(testRepo)
	testHandler := user.NewHandler(testService)
	testHandler.RegisterRoutes(api)

	return r
}
