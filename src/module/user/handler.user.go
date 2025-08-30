package user

import (
	"gin/src/config"
	"gin/src/share"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
			h.GetUser,
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
			"GET",
			"/validate",
			share.LogRequest(logger),
			share.AuthMiddleware(),
			h.Validate,
		)

		userGroup.Handle(
			"GET",
			"/.well-known/jwks.json",
			share.LogRequest(logger),
			h.JWK,
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

		userGroup.Handle(
			"PATCH",
			"/delete",
			share.LogRequest(logger),
			share.RateLimitMiddleware(config.GlobalRatelimit),
			share.AuthMiddleware(),
			h.PatchDeleteUser,
		)

	}
}
