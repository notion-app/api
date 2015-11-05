package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	currentStatus = map[string]interface{}{
		"status": 200,
	}
)

// Status handler for the /status route
func Status(c *gin.Context) {
	c.JSON(http.StatusOK, currentStatus)
}
