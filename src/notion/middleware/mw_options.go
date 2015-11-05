package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AcceptOptions(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
	}
}
