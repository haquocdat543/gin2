package user

import (
	"gin/src/share"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) CreateUser(
	c *gin.Context,
) {
	var dto CreateUserDTO

	if !share.BindJSONAndValidate(c, &dto) {
		return // the function already handled the error response
	}

	// Convert DTO to Entity
	user := User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}

	// Error handle
	err := h.service.CreateUser(&user)
	if err != nil {
		if strings.Contains(
			err.Error(),
			"ERROR: duplicate key value violates unique constraint \"idx_users_email\" (SQLSTATE 23505)",
		) ||
			strings.Contains(
				err.Error(),
				"UNIQUE constraint failed",
			) {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": ErrEmailAlreadyExists,
				},
			)
		} else {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": ErrInternalServer,
				},
			)
		}
		return
	}

	// Data return
	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": MsgUserCreated,
			"user":    user,
		},
	)
}
