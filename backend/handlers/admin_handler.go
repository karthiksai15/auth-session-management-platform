package handlers

import (
	"auth-system/backend/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllUsers handles GET /admin/users
// Only reachable by ADMIN role — enforced by RoleMiddleware in the routes.
func GetAllUsers(c *gin.Context) {
	users, err := repository.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
