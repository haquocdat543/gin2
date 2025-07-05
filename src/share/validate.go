package share

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// BindAndValidate binds the JSON body to the given DTO and validates it.
// Returns false if there's an error, and sends the appropriate response.
func BindAndValidate[T any](c *gin.Context, dto *T) bool {
	if err := c.ShouldBindJSON(dto); err != nil {
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
		return false
	}
	return true
}

