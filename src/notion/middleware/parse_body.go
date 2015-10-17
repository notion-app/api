
package middleware

import (
  "fmt"
  "encoding/json"
  "github.com/labstack/echo"
  "io/ioutil"
  "notion/errors"
  "notion/log"
)

func ParseBody() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {

      var request map[string]interface{}
      b, err := ioutil.ReadAll(c.Request().Body)
    	if log.Error(err) {
    		return errors.BadRequest("Error reading request body")
    	}
    	_ = json.Unmarshal(b, &request)

      fmt.Printf("%v\n", request)

      c.Set("BODY", request)

      if err := h(c); err != nil {
				c.Error(err)
			}
			return nil

    }
  }
}
