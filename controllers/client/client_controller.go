// controllers/client/client_controller.go
package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ClientController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ClientController dashboard"})
}
