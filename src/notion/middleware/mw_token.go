package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"notion/db"
	// "notion/errors"
	"notion/log"
)

func AuthCheck(c *gin.Context) {
	id, _ := c.Get("request_id")
	log.InfoFields("Checking auth token", log.Fields{
		"request_id": id,
	})

	token := c.Query("token")
	if token == "" {
		tokenHeaders := c.Request.Header["Token"]
		if len(tokenHeaders) == 0 {
			c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("No notion access token provided"))
			return
		} else {
			token = tokenHeaders[0]
		}
	}

	in, user, err := db.GetUserByToken(token)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Internal server error"))
		return
	}
	if !in {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Notion access token provided is not currently valid"))
		return
	}

	c.Set("request_user_id", user.Id)
}
