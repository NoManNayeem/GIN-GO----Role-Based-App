// middlewares/cors_middleware.go
package middlewares

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func CORSMiddleware() gin.HandlerFunc {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	return func(c *gin.Context) {
		// Read environment variables from .env file
		allowOrigin := os.Getenv("CORS_ALLOW_ORIGIN")
		allowCredentials := os.Getenv("CORS_ALLOW_CREDENTIALS")
		allowHeaders := os.Getenv("CORS_ALLOW_HEADERS")
		allowMethods := os.Getenv("CORS_ALLOW_METHODS")

		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", allowCredentials)
		c.Writer.Header().Set("Access-Control-Allow-Headers", allowHeaders)
		c.Writer.Header().Set("Access-Control-Allow-Methods", allowMethods)

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
