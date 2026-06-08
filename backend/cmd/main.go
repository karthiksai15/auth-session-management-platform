package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router with default middleware (logger + recovery)
	r := gin.Default()

	// Health check endpoint — used to verify the server is running
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Start the server on port 8080
	fmt.Println("Server running on port 8080")
	r.Run(":8080")
}
