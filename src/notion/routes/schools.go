
package routes

import (
  "github.com/labstack/echo"
  "net/http"
  "notion/db"
  "notion/errors"
  "notion/model"
)

func GetSchools(c *echo.Context) error {
  schools, err := db.GetAllSchools()
  if err != nil {
    return errors.ISE()
  }
  return c.JSON(http.StatusOK, model.AllSchoolsResponse{
    Schools: schools,
  })
}

func PostSchoolRequest(c *echo.Context) error {
  return nil
}
