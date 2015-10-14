
package routes

import (
  "encoding/json"
  "github.com/labstack/echo"
  "io/ioutil"
  "net/http"
  "notion/db"
  "notion/errors"
  "notion/log"
  "notion/model"
  "notion/validate"
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
  var request model.SchoolRequestRequest
  b, err := ioutil.ReadAll(c.Request().Body)
  if log.Error(err) {
    return errors.BadRequest("Error reading request body")
  }
  err = json.Unmarshal(b, &request)
  if log.Error(err) {
    return errors.BadRequest("Error parsing json body")
  }
  err = validate.SchoolRequest(request)
  if log.Error(err) {
    return err
  }
  return nil
}
