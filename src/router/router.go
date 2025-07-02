package router

import (
	test "gin/src/module/test"
	test2 "gin/src/module/test2"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		test.RegisterRoutes(api)
		test2.RegisterRoutes(api)
	}

	return r
}
