package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

var (
	currentStatus = map[string]interface{}{
		"status": 200,
	}
)

// Status handler for the /status route
func Status(c *echo.Context) error {
	return c.JSON(http.StatusOK, currentStatus)
}
