package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(
	rg *gin.RouterGroup,
) {
	test := rg.Group("/test")
	{

		test.Handle(
			"GET",
			"/ping",
			PingPong,
		)

		test.Handle(
			"GET",
			"/ask",
			AskAnswer,
		)

	}
}

func PingPong(
	c *gin.Context,
) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "pong",
		},
	)
}

func AskAnswer(
	c *gin.Context,
) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "answer",
		},
	)
}
