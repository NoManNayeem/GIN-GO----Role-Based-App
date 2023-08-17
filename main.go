package main

import (
	"role_based_app/initializers"
	"role_based_app/middlewares"
	"role_based_app/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func main() {
	router := gin.Default()

	// CORS
	router.Use(middlewares.CORSMiddleware())
	// Routers
	routes.SetupRoutes(router)
	router.Run()
}
