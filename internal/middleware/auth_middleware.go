package middleware

import (
	"net/http"

	"github.com/ardianilyas/go-ticketing/internal/jwt"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT token from cookie and attaches user to context
func AuthMiddleware(jwtService *jwt.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from cookie
		token, err := c.Cookie("access_token")
		if err != nil || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": map[string]interface{}{
					"code":    "UNAUTHORIZED",
					"message": "Authentication required",
				},
			})
			c.Abort()
			return
		}

		// Validate token
		claims, err := jwtService.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": map[string]interface{}{
					"code":    "UNAUTHORIZED",
					"message": "Invalid or expired token",
				},
			})
			c.Abort()
			return
		}

		// Set user info in context
		c.Set("user", map[string]interface{}{
			"id":    claims.UserID.String(),
			"email": claims.Email,
		})
		c.Set("userId", claims.UserID.String())

		c.Next()
	}
}

// OptionalAuth is middleware that extracts user from token if present, but doesn't require it
func OptionalAuth(jwtService *jwt.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("access_token")
		if err == nil && token != "" {
			claims, err := jwtService.ValidateToken(token)
			if err == nil {
				c.Set("user", map[string]interface{}{
					"id":    claims.UserID.String(),
					"email": claims.Email,
				})
				c.Set("userId", claims.UserID.String())
			}
		}
		c.Next()
	}
}
