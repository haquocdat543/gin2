package user

import (
	"gin/pkg/share"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) UpdatePassword(
	c *gin.Context,
) {
	var dto UpdatePasswordDTO

	// Bind JSON to DTO and validate
	if !share.BindJSONAndValidate(c, &dto) {
		return // the function already handled the error response
	}

	// Error handle
	err := h.service.Login(dto.Name, dto.Password)
	if err != nil {
		c.JSON(
			400,
			gin.H{
				"error": err.Error(),
			},
		)
	} else {

		token, err := share.GenerateToken(dto.Name, c.ClientIP())
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": "Could not generate token",
				},
			)
		} else {

			err := h.service.UpdateUserPassword(dto.Name, dto.NewPassword)
			if err != nil {
				c.JSON(
					http.StatusInternalServerError,
					gin.H{
						"error": "Failed to update password",
					},
				)
			}

		}

		// Data return
		c.JSON(
			http.StatusCreated,
			gin.H{
				"message": MsgLoginSuccess,
				"jwt":     token,
			},
		)
	}

}
