package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"validate": "OK",
	})
}
