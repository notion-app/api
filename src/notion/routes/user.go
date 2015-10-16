package routes

import (
	"github.com/labstack/echo"
	"net/http"
	"notion/db"
	"notion/errors"
)

func GetUser(c *echo.Context) error {
	userId := c.Param("user_id")
	in, user, err := db.GetUserById(userId)
	if err != nil {
		return errors.ISE()
	}
	if !in {
		return errors.NotFound()
	}
	return c.JSON(http.StatusOK, user)
}
