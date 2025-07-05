package user

import (
	"gin/src/share"
	"github.com/gin-gonic/gin"
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

	if !share.BindAndValidate(c, &dto) {
		return // the function already handled the error response
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

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": MsgUserCreated,
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
	if !share.BindAndValidate(c, &dto) {
		return // the function already handled the error response
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

		token, err := share.GenerateToken(dto.Name)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": "Could not generate token",
				},
			)
			return
		}

		c.JSON(
			http.StatusCreated,
			gin.H{
				"message": MsgLoginSuccess,
				"jwt":     token,
			},
		)
	}

}
