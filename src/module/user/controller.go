package user

import (
	"errors"
	"fmt"
	"gin/src/config"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type Handler struct {
	service Service
}

func NewHandler(
	s Service,
) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) RegisterRoutes(
	rg *gin.RouterGroup,
) {
	userGroup := rg.Group("/user")
	{

		userGroup.Handle(
			"POST",
			"/",
			h.CreateUser,
		)

		userGroup.Handle(
			"GET",
			"/",
			h.GetUsers,
		)

	}
}

func (h *Handler) CreateUser(
	c *gin.Context,
) {
	var dto CreateUserDTO

	// Bind JSON to DTO and validate
	if err := c.ShouldBindJSON(&dto); err != nil {
		// Parse validation errors
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make(map[string]string)
			for _, fe := range ve {
				out[fe.Field()] = fmt.Sprintf("failed on '%s' tag", fe.Tag())
			}
			c.JSON(400, gin.H{"errors": out})
			return
		}
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Manually map DTO to Entity
	user := User{
		Name:  dto.Name,
		Email: dto.Email,
		Age:   uint(dto.Age), // safe conversion
	}

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
					"error": config.ErrEmailAlreadyExists,
				},
			)
		} else {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": config.ErrInternalServer,
				},
			)
		}
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": config.MsgUserCreated,
			"user":    user,
		},
	)
}

func (h *Handler) GetUsers(
	c *gin.Context,
) {

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
	c.JSON(
		http.StatusOK,
		users,
	)
}
