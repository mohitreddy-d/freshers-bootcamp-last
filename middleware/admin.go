package middleware

import (
	"freshers-bootcamp-last/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

var secretKey = os.Getenv("JWT_SECRET")

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Validate the Bearer token format
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// Parse and validate the JWT token
		claims, _, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		role := claims.Role
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		// Proceed to the next middleware or handler
		c.Next()
	}
}
