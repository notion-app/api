
package errors

import (
  "github.com/labstack/echo"
  "net/http"
)

func BadRequest(c *echo.Context, msg string) error {
  c.JSON(http.StatusBadRequest, map[string]interface{}{
    "status": http.StatusBadRequest,
    "message": msg,
  })
  return nil
}

// Alerts the client that an access token for a given service is expired.
// 'Service' should usually be one of 'facebook' or 'notion'.
func Unauthorized(c *echo.Context, service string) error {
  c.JSON(http.StatusUnauthorized, map[string]interface{}{
    "status": http.StatusUnauthorized,
    "service": service,
    "message": "The access token you have provided does not have permission to view the resource you've requested, or has expired.",
  })
  return nil
}

func NotFound(c *echo.Context) error {
  c.JSON(http.StatusNotFound, map[string]interface{}{
    "status": http.StatusNotFound,
    "message": "Requested resource could not be found",
  })
  return nil
}

func ISE(c *echo.Context) error {
  c.JSON(http.StatusInternalServerError, map[string]interface{}{
    "status": http.StatusInternalServerError,
    "message": "Internal Server Error",
  })
  return nil
}
