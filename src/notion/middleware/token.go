package middleware

import (
	"fmt"
	"github.com/labstack/echo"
	"notion/logic"
)

func TokenCheck() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {

			token := c.Query("token")
			if token == "" {
				return fmt.Errorf("Unauthorized")
			}

			user, err := logic.AuthenticateNotionUser(token)
			if err != nil {
				return err
			}

			c.Set("TOKEN_USER_ID", user.Id)

			if err := h(c); err != nil {
				c.Error(err)
			}

			return nil

		}
	}
}
