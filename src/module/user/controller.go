package user

import (
	"gin/src/config"
	"github.com/gin-gonic/gin"
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
	if err := c.ShouldBindJSON(
		&dto,
	); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	// Validating
	if err := dto.Validate(); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"errors": err,
			},
		)
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
