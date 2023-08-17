// middlewares/moderator_role_required.go
package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ModeratorRoleRequired(c *gin.Context) {
	// Retrieve role information from the context
	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Role not found in context",
		})
		c.Abort()
		return
	}
	// Check if the user has the "moderator" role
	if role != "moderator" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Access denied. Moderator role required.",
		})
		c.Abort()
		return
	}

	// Continue to the next middleware or handler
	c.Next()
}
