package handlers

import (
	"auth-system/backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRequest defines the expected JSON body for POST /auth/register
type RegisterRequest struct {
	Name     string `json:"name"     binding:"required"`
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// Register handles POST /auth/register
// It reads name, email, and password from the request body,
// then calls the service to create the user.
func Register(c *gin.Context) {
	var req RegisterRequest

	// Bind the JSON body to RegisterRequest and validate it
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the auth service to register the user
	user, err := services.RegisterUser(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the created user (without the password)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}

// LoginRequest defines the expected JSON body for POST /auth/login
type LoginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Login handles POST /auth/login
// It checks the email and password, then returns an access token and a refresh token.
func Login(c *gin.Context) {
	var req LoginRequest

	// Bind and validate the request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the auth service — it checks credentials and generates tokens
	accessToken, refreshToken, err := services.LoginUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Return both tokens to the client
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// RefreshRequest defines the expected JSON body for POST /auth/refresh
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// Refresh handles POST /auth/refresh
// Takes a refresh token, validates it, checks Redis, and returns a new access token.
func Refresh(c *gin.Context) {
	var req RefreshRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newAccessToken, err := services.RefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
	})
}

// Logout handles POST /auth/logout
// AuthMiddleware runs first and stores userId in context.
// This handler reads that userId and deletes the refresh token from Redis.
func Logout(c *gin.Context) {
	// Get the userId from Gin context — AuthMiddleware put it there
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Call the service to delete the refresh token from Redis
	err := services.LogoutUser(userId.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
