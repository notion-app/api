package middleware

import (
	"github.com/gin-gonic/gin"
	"notion/log"
	"notion/util"
	"time"
)

func Logger(c *gin.Context) {
	requestId := util.NewId()
	c.Set("request_id", requestId)

	method := c.Request.Method
	path := c.Request.URL.EscapedPath()
	ip := c.ClientIP()

	log.InfoFields("Request received", log.Fields{
		"request_id": requestId,
		"method":     method,
		"ip":         ip,
		"path":       path,
	})

	start := time.Now()
	c.Next()
	duration := time.Since(start)

	code := c.Writer.Status()

	log.InfoFields("Request handled", log.Fields{
		"request_id": requestId,
		"took":       duration.String(),
		"code":       code,
	})

}
