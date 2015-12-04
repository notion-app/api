package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

var (
	limiter = time.Tick(time.Millisecond * 10)
)

func RateLimit(c *gin.Context) {
	<-limiter
}
