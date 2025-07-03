package router

import (
	"gorm.io/gorm"
	"gin/src/module/test"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	// ✅ Initialize service and handler properly
	testRepo := test.NewRepository(db)
	testService := test.NewService(testRepo)
	testHandler := test.NewHandler(testService)
	testHandler.RegisterRoutes(api)

	return r
}
