// controllers/admin/admin_controller.go
package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Admin dashboard"})
}
