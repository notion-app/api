
package middleware

import (
	"github.com/labstack/echo"
  "notion/errors"
  "notion/logic"
)

func TokenCheck() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {

      token := c.Query("token")
      if token == "" {
        return errors.Unauthorized("notion")
      }

      user, err := logic.AuthenticateNotionUser(token)
      if err != nil {
        return err
      }

      c.Set("TOKEN_USER_ID", user.Id)
      return nil

    }
  }
}
