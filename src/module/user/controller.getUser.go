package user

import (
	"gin/src/share"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Find(
	c *gin.Context,
) {

	username, err := share.GetUsername(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Find(username)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
	}

	// Data return
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": MsgUserInfoFetched,
			"user":    user,
		},
	)
}
