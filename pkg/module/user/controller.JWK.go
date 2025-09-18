package user

import (
	"gin/pkg/share"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) JWK(c *gin.Context) {
	jwkSet := share.CreateJWKSet()
	c.JSON(http.StatusOK, jwkSet)
}
