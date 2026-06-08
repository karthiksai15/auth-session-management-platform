package handlers

import (
	"auth-system/backend/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProfile handles GET /users/profile
// Requires valid JWT token via AuthMiddleware.
func GetProfile(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	user, err := repository.FindUserByID(userId.(int))
	if err != nil || user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// UpdateProfileRequest defines the allowed fields for profile update
type UpdateProfileRequest struct {
	Name string `json:"name" binding:"required"`
}

// UpdateProfile handles PUT /users/profile
// Only allows updating the name.
func UpdateProfile(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repository.UpdateUser(userId.(int), req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}
