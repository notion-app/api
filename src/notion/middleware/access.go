
package middleware

import (
  "github.com/labstack/echo"
)

func AccessControl() echo.MiddlewareFunc {
  return func(h echo.HandlerFunc) echo.HandlerFunc {
    return func(c *echo.Context) error {

      if err := h(c); err != nil {
				c.Error(err)
			}

      if origin := c.Request().Header.Get("Origin"); origin != "" {
        c.Response().Header().Set("Access-Control-Allow-Origin", origin)
      }
      c.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
      c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
      c.Response().Header().Set("Access-Control-Allow-Credentials", "true")

      return nil

    }
  }
}