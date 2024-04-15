package middlewares

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"Brilliant/application/util"
)

// AuthMiddleware extracts the bearer token from the request header and sets the user details to the context.
func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken := c.GetHeader("RefreshToken")
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		claims, err := util.ExtractUserDetails(tokenString, secret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Set("refresh_token", refreshToken)

		c.Next()
	}
}

// RoleMiddleware checks if the user in the context has the specified role.
func RoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, _ := c.Get("role")
		if userRole != role {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		c.Next()
	}
}