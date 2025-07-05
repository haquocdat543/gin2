package user

import (
	"gin/src/config"
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
			share.RateLimitMiddleware(config.GlobalRatelimit),
			h.CreateUser,
		)

		userGroup.Handle(
			"GET",
			"/",
			share.LogRequest(logger),
			share.RateLimitMiddleware(config.GlobalRatelimit),
			share.AuthMiddleware(),
			h.Find,
		)

		userGroup.Handle(
			"DELETE",
			"/",
			share.LogRequest(logger),
			share.RateLimitMiddleware(config.GlobalRatelimit),
			h.DeleteUser,
		)

		userGroup.Handle(
			"GET",
			"/all",
			share.LogRequest(logger),
			share.RateLimitMiddleware(config.GlobalRatelimit),
			h.GetUsers,
		)

		userGroup.Handle(
			"POST",
			"/login",
			share.LogRequest(logger),
			share.RateLimitMiddleware(config.GlobalRatelimit),
			h.Login,
		)

		userGroup.Handle(
			"PATCH",
			"/password",
			share.LogRequest(logger),
			share.RateLimitMiddleware(config.GlobalRatelimit),
			h.UpdatePassword,
		)

		userGroup.Handle(
			"PATCH",
			"/",
			share.LogRequest(logger),
			share.RateLimitMiddleware(config.GlobalRatelimit),
			share.AuthMiddleware(),
			h.PatchUpdateUser,
		)

		userGroup.Handle(
			"PUT",
			"/",
			share.LogRequest(logger),
			share.RateLimitMiddleware(config.GlobalRatelimit),
			share.AuthMiddleware(),
			h.PutUpdateUser,
		)
	}
}

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

func (h *Handler) Find(
	c *gin.Context,
) {

	username, err := share.GetUsername(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Find(username)
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
			"message": MsgUserInfoFetched,
			"user":    user,
		},
	)
}

func (h *Handler) GetUsers(
	c *gin.Context,
) {

	// Error handle
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

	// Data return
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

		// Data return
		c.JSON(
			http.StatusCreated,
			gin.H{
				"message": MsgLoginSuccess,
				"jwt":     token,
			},
		)
	}

}

func (h *Handler) UpdatePassword(
	c *gin.Context,
) {
	var dto UpdatePasswordDTO

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

		token, err := share.GenerateToken(dto.Name)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": "Could not generate token",
				},
			)
		} else {

			err := h.service.UpdateUserPassword(dto.Name, dto.NewPassword)
			if err != nil {
				c.JSON(
					http.StatusInternalServerError,
					gin.H{
						"error": "Failed to update password",
					},
				)
			}

		}

		// Data return
		c.JSON(
			http.StatusCreated,
			gin.H{
				"message": MsgLoginSuccess,
				"jwt":     token,
			},
		)
	}

}

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

func (h *Handler) PutUpdateUser(
	c *gin.Context,
) {
	var dto PutUserDTO

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
	user.Dob = share.ParseDate(dto.Dob)
	user.Role = &dto.Role
	user.Address = &dto.Address

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
