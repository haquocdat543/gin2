package share

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// RateLimitMiddleware creates a rate-limiting middleware for a specific rate.
func RateLimitMiddleware(
	rate limiter.Rate,
) gin.HandlerFunc {
	// Create a memory store for the rate limiter.
	store := memory.NewStore()

	// Create a new rate limiter instance.
	instance := limiter.New(
		store,
		rate,
	)

	return func(
		c *gin.Context,
	) {
		// Use the client's IP address as the key for rate limiting.
		ip := c.ClientIP()

		// Check the rate limit for the client.
		context, err := instance.Get(
			c,
			ip,
		)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{
					"error": "Internal Server Error",
				},
			)
			return
		}

		// Set rate limit headers in the response.
		c.Header(
			"X-RateLimit-Limit",
			"X-RateLimit-Limit",
		)
		c.Header(
			"X-RateLimit-Remaining",
			"X-RateLimit-Remaining",
		)
		c.Header(
			"X-RateLimit-Reset",
			"X-RateLimit-Reset",
		)

		// If the client has exceeded the rate limit, return a 429 Too Many Requests response.
		if context.Reached {
			c.AbortWithStatusJSON(
				http.StatusTooManyRequests,
				gin.H{
					"error": "Too Many Requests",
				},
			)
			return
		}

		// Continue to the next handler if the rate limit is not exceeded.
		c.Next()
	}
}
