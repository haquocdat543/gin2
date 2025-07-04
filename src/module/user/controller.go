package user

import (
	"fmt"
	"gin/src/config"
	"gin/src/share"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
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
	logger *zap.Logger,
) {
	userGroup := rg.Group("/user")
	{

		userGroup.Handle(
			"POST",
			"/",
			share.LogRequest(logger),
			share.RateLimitMiddleware(share.GlobalRatelimit),
			h.CreateUser,
		)

		userGroup.Handle(
			"GET",
			"/",
			share.LogRequest(logger),
			share.RateLimitMiddleware(share.GlobalRatelimit),
			h.GetUsers,
		)

		userGroup.Handle(
			"POST",
			"/login",
			share.LogRequest(logger),
			share.RateLimitMiddleware(share.GlobalRatelimit),
			h.Login,
		)

	}
}

func (h *Handler) CreateUser(
	c *gin.Context,
) {
	var dto CreateUserDTO

	// Bind JSON to DTO and validate
	if err := c.ShouldBindJSON(&dto); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			errorsMap := make(map[string]string, len(ve))
			for _, fe := range ve {
				errorsMap[fe.Field()] = fmt.Sprintf(
					"Field '%s' failed validation: tag='%s', param='%s'",
					fe.Field(), fe.Tag(), fe.Param(),
				)
			}
			c.JSON(400, gin.H{"errors": errorsMap})
		} else {
			c.JSON(400, gin.H{"error": err.Error()})
		}
		return
	}

	// Manually map DTO to Entity
	user := User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
		Age:      uint(dto.Age), // safe conversion
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

func (h *Handler) Login(
	c *gin.Context,
) {
	var dto LoginDTO

	// Bind JSON to DTO and validate
	if err := c.ShouldBindJSON(&dto); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			errorsMap := make(map[string]string, len(ve))
			for _, fe := range ve {
				errorsMap[fe.Field()] = fmt.Sprintf(
					"Field '%s' failed validation: tag='%s', param='%s'",
					fe.Field(), fe.Tag(), fe.Param(),
				)
			}
			c.JSON(400, gin.H{"errors": errorsMap})
		} else {
			c.JSON(400, gin.H{"error": err.Error()})
		}
		return
	}

	err := h.service.Login(dto.Name, dto.Password)
	if err != nil {
		c.JSON(
			400,
			gin.H{
				"error": err.Error(),
			},
		)
	} else {
		c.JSON(
			http.StatusCreated,
			gin.H{
				"message": config.MsgLoginSuccess,
			},
		)
	}

}
