
package errors

import (
  "github.com/labstack/echo"
  "net/http"
)

func BadRequest(msg string) error {
  return echo.NewHTTPError(http.StatusBadRequest, msg)
}

// Alerts the client that an access token for a given service is expired.
// 'Service' should usually be one of 'facebook' or 'notion'.
func Unauthorized(service string) error {
  return echo.NewHTTPError(http.StatusUnauthorized, "The access token you have provided does not have permission to view the resource you've requested, or has expired.")
}

func NotFound() error {
  return echo.NewHTTPError(http.StatusNotFound, "Requested resource could not be found")
}

func ISE() error {
  return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
}
