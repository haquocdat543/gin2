package user

import (
	"gin/src/share"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) DeleteUser(
	c *gin.Context,
) {
	var dto DeleteUserDTO

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

		err := h.service.DeleteUser(dto.Name)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": "Failed to update password",
				},
			)
		}

		// Data return
		c.JSON(
			http.StatusCreated,
			gin.H{
				"message": MsgDeleteSuccess,
			},
		)
	}

}

