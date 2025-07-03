package test

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
	userGroup := rg.Group("/test")
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

	var user Test

	if err := c.ShouldBindJSON(
		&user,
	); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
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

