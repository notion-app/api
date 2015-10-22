package middleware

import (
	"github.com/labstack/echo"
)

func AccessControl() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {

			// Lol dat security do
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")	
			c.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
			c.Response().Header().Set("Access-Control-Allow-Headers", "*")
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")

			if err := h(c); err != nil {
				c.Error(err)
			}
			return nil

		}
	}
}
