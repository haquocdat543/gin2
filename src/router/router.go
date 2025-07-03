package router

import (
	test "gin/src/module/test"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		test.RegisterRoutes(api)
	}

	return r
}
