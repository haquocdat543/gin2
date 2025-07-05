package share

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// handleValidationErrors converts validator.ValidationErrors to a map and responds with 400.
func handleValidationErrors(c *gin.Context, ve validator.ValidationErrors) {
	errorsMap := make(map[string]string, len(ve))
	for _, fe := range ve {
		errorsMap[fe.Field()] = fmt.Sprintf(
			"Field '%s' failed validation: tag='%s', param='%s'",
			fe.Field(), fe.Tag(), fe.Param(),
		)
	}
	c.JSON(http.StatusBadRequest, gin.H{"errors": errorsMap})
}

// BindJSONAndValidate binds JSON body to dto and validates it.
func BindJSONAndValidate[T any](c *gin.Context, dto *T) bool {
	err := c.ShouldBindJSON(dto)
	if err != nil {
		if err == io.EOF {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Request body is empty"})
			return false
		}
		if ve, ok := err.(validator.ValidationErrors); ok {
			handleValidationErrors(c, ve)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return false
	}
	return true
}

// BindQueryAndValidate binds query parameters to dto and validates them.
func BindQueryAndValidate[T any](c *gin.Context, dto *T) bool {
	if err := c.ShouldBindQuery(dto); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			handleValidationErrors(c, ve)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return false
	}
	return true
}

// BindUriAndValidate binds path parameters to dto and validates them.
func BindUriAndValidate[T any](c *gin.Context, dto *T) bool {
	if err := c.ShouldBindUri(dto); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			handleValidationErrors(c, ve)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return false
	}
	return true
}

