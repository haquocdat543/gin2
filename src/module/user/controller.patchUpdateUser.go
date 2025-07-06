package user

import (
	"gin/src/share"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) PatchUpdateUser(
	c *gin.Context,
) {
	var dto PatchUserDTO

	// Bind JSON to DTO and validate
	if !share.BindJSONAndValidate(c, &dto) {
		return // the function already handled the error response
	}

	username, err := share.GetUsername(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	user := User{}

	user.Name = username

	if dto.Dob != nil {
		user.Dob = share.ParseDate(*dto.Dob)
	}

	if dto.Role != nil {
		user.Role = dto.Role
	}

	if dto.Address != nil {
		user.Address = dto.Address
	}

	err = h.service.UpdateUser(&user)
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
			"message": MsgUpdateUserSuccess,
		},
	)

}

