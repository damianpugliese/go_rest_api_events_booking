package middlewares

import (
	"net/http"

	"github.com/damianpugliese/go_rest_api_events_booking/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Authorization token is required",
		})
		return  
	}

	userId, err := utils.ValidateJWT(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Invalid or expired token",
			"error":   err.Error(),
		})
		return
	}

	c.Set("userId", userId)

	c.Next()
}