package main

import (
	"fmt"
	"log"
	"net/http"

	"auth-system/backend/config"
	"auth-system/backend/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to PostgreSQL
	config.ConnectDB()

	// Connect to Redis
	config.ConnectRedis()

	// Create a new Gin router with default middleware (logger + recovery)
	r := gin.Default()

	// Health check endpoint — used to verify the server is running
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Auth routes — public (no authentication required)
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	// Start the server on port 8080
	fmt.Println("Server running on port 8080")
	r.Run(":8080")
}
