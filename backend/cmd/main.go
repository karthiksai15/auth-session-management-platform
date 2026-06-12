package main

import (
	"fmt"
	"net/http"

	"auth-system/backend/config"
	"auth-system/backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	config.ConnectDB()
	config.ConnectRedis()

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	routes.SetupRoutes(r)

	fmt.Println("Server running on port 8080")
	r.Run(":8080")
}
