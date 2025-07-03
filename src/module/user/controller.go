package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
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

	if err := h.service.CreateUser(
		&user,
	); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		user,
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
