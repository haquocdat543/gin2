package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetUsers(
	c *gin.Context,
) {

	// Error handle
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	// Data return
	c.JSON(
		http.StatusOK,
		users,
	)
}
