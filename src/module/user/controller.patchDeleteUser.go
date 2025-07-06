package user

import (
	"gin/src/share"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) PatchDeleteUser(
	c *gin.Context,
) {
	var dto PatchDeleteDTO

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

	fields := make(map[string]any)
	if dto.Dob != nil && *dto.Dob {
		fields["dob"] = nil
	}
	if dto.Role != nil && *dto.Role {
		fields["role"] = nil
	}
	if dto.Address != nil && *dto.Address {
		fields["address"] = nil
	}

	err = h.service.PatchDeleteUser(&user, fields)
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
			"message": MsgPatchDeleteUserSuccess,
		},
	)

}
