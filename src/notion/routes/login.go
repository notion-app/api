package routes

import (
  "encoding/json"
  "github.com/labstack/echo"
  "io/ioutil"
  "notion/errors"
  "notion/log"
  "notion/logic"
  "notion/model"
  "notion/validate"
)

func Login(c *echo.Context) error {
  var request model.LoginRequest
  b, err := ioutil.ReadAll(c.Request().Body)
  if log.Error(err) {
    return errors.BadRequest("Error reading request body")
  }
  err = json.Unmarshal(b, &request)
  if log.Error(err) {
    return errors.BadRequest("Error parsing json body")
  }
  err = validate.Login(request)
  if log.Error(err) {
    return err
  }
  code, resp, err := logic.DoUserCreateOrLogin(request)
  if log.Error(err) {
    return err
  }
  return c.JSON(code, resp)
}
