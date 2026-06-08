package middleware

import (
	"auth-system/backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates the JWT token from the Authorization header.
// If the token is valid, it stores userId and role in the Gin context.
// If the token is missing or invalid, it returns 401 Unauthorized and stops the request.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Step 1: Read the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort() // Stop processing — don't call the next handler
			return
		}

		// Step 2: The header should look like "Bearer <token>"
		// Split on the space to get the two parts
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid format. Use: Bearer <token>"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Step 3: Validate the token using our JWT utility
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Step 4: Store userId and role in the Gin context
		// Handlers downstream can read these with c.Get("userId") and c.Get("role")
		c.Set("userId", claims.UserID)
		c.Set("role", claims.Role)

		// Token is valid — move on to the next middleware or handler
		c.Next()
	}
}

// RoleMiddleware checks if the authenticated user's role matches the required role.
// This must run AFTER AuthMiddleware (which stores the role in the context).
// Returns 403 Forbidden if the role does not match.
func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read the role from the Gin context (set by AuthMiddleware)
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Role not found in context"})
			c.Abort()
			return
		}

		// Simple string comparison — one role check, nothing more
		if role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: insufficient role"})
			c.Abort()
			return
		}

		// Role matches — continue to the handler
		c.Next()
	}
}
