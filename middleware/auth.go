package middleware

import (
	"fmt"
	"freshers-bootcamp-last/utils"
	"github.com/gin-gonic/gin"

	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		fmt.Println(token)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization token required"})
			c.Abort()
			return
		}

		claims, _, err := utils.ValidateToken(token[7:])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
