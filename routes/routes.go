package routes

import (
	"role_based_app/controllers/admin"
	"role_based_app/controllers/auth"
	"role_based_app/controllers/client"
	"role_based_app/controllers/moderator"
	"role_based_app/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// Auth routes
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", auth.Login)
		authGroup.POST("/signup", auth.Signup)
	}

	// Protected routes requiring authentication
	protectedGroup := r.Group("/")
	protectedGroup.Use(middlewares.LoginRequired)
	{
		// Admin routes
		adminGroup := protectedGroup.Group("/admin")
		adminGroup.Use(middlewares.AdminRoleRequired)
		{
			adminGroup.GET("/dashboard", admin.AdminController)
		}

		// Moderator routes
		moderatorGroup := protectedGroup.Group("/moderator")
		moderatorGroup.Use(middlewares.ModeratorRoleRequired)
		{
			moderatorGroup.GET("/dashboard", moderator.ModeratorController)
		}

		// Client routes
		clientGroup := protectedGroup.Group("/client")
		clientGroup.Use(middlewares.ClientRoleRequired)
		{
			clientGroup.GET("/dashboard", client.ClientController)
		}
	}
}
