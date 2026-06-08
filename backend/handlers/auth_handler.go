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
