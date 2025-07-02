package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	test := r.Group("/test")
	{
		test.Handle(
			"GET",
			"/ping",
			func(c *gin.Context) {
				c.JSON(
					http.StatusOK,
					gin.H{
						"message": "pong",
					},
				)
			},
		)

		test.Handle(
			"GET",
			"/ask",
			func(c *gin.Context) {
				c.JSON(
					http.StatusOK,
					gin.H{
						"message": "answer",
					},
				)
			},
		)
	}

	test2 := r.Group("/test2")
	{
		test2.Handle(
			"GET",
			"/ping",
			func(c *gin.Context) {
				c.JSON(
					http.StatusOK,
					gin.H{
						"message": "pong2",
					},
				)
			},
		)

		test2.Handle(
			"GET",
			"/ask",
			func(c *gin.Context) {
				c.JSON(
					http.StatusOK,
					gin.H{
						"message": "answer2",
					},
				)
			},
		)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
