package middleware

import (
	"github.com/gin-gonic/gin"
	"notion/log"
)

func AccessControl(c *gin.Context) {
	id, _ := c.Get("request_id")
	log.InfoFields("Setting access control headers", log.Fields{
		"request_id": id,
	})
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
}
