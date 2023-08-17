// controllers/moderator/moderator_controller.go
package moderator

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ModeratorController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ModeratorController dashboard"})
}
