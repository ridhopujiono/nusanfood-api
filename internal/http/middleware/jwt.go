package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ridhopujiono/nusanfood-api/internal/auth"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Missing or invalid Authorization header",
			})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := auth.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid or expired token",
			})
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
