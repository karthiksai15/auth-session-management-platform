package routes

import (
	"auth-system/backend/handlers"
	"auth-system/backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
		auth.POST("/refresh", handlers.Refresh)
		auth.POST("/logout", middleware.AuthMiddleware(), handlers.Logout)
	}

	users := r.Group("/users")
	users.Use(middleware.AuthMiddleware())
	{
		users.GET("/profile", handlers.GetProfile)
		users.PUT("/profile", handlers.UpdateProfile)
	}

	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.RoleMiddleware("ADMIN"))
	{
		admin.GET("/users", handlers.GetAllUsers)
	}
}
