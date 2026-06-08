package main

import (
	"fmt"
	"log"
	"net/http"

	"auth-system/backend/config"
	"auth-system/backend/handlers"
	"auth-system/backend/middleware"

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
		auth.POST("/refresh", handlers.Refresh)
		// Logout requires a valid token — user must be logged in
		auth.POST("/logout", middleware.AuthMiddleware(), handlers.Logout)
	}

	// User routes — protected (valid JWT required)
	// Full handlers will be added in Phase 10
	users := r.Group("/users")
	users.Use(middleware.AuthMiddleware())
	{
		users.GET("/profile", func(c *gin.Context) {
			// Temporary placeholder — replaced by handlers.GetProfile in Phase 10
			userId, _ := c.Get("userId")
			role, _ := c.Get("role")
			c.JSON(http.StatusOK, gin.H{
				"message": "Auth working — full profile coming in Phase 10",
				"userId":  userId,
				"role":    role,
			})
		})
		users.PUT("/profile", func(c *gin.Context) {
			// Temporary placeholder — replaced by handlers.UpdateProfile in Phase 10
			c.JSON(http.StatusOK, gin.H{
				"message": "Auth working — full update coming in Phase 10",
			})
		})
	}

	// Admin routes — require valid JWT AND role must be ADMIN
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.RoleMiddleware("ADMIN"))
	{
		admin.GET("/users", handlers.GetAllUsers)
	}

	// Start the server on port 8080
	fmt.Println("Server running on port 8080")
	r.Run(":8080")
}

