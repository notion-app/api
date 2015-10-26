
// This middleware handles transposing any error which is returned by
// the notion application into a json-parsed machine readable error

package middleware

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "notion/errors"
  "notion/log"
  "strings"
)

func Error(c *gin.Context) {
  c.Next()
  id, _ := c.Get("request_id")

  // Log out every error we have encoutered (which in most cases is just 1)
  for _, ginError := range c.Errors {
    actError := ginError.Err
    log.InfoFields("Request error", log.Fields{
      "request_id": id,
      "body": formatErrorBody(actError.Error()),
    })
  }

  // Grab the last error and use that as the error we return to the client
  if len(c.Errors) > 0 {
    clientError := c.Errors[len(c.Errors)-1].Err

    // If it isn't an errors.Http type, assume it is a 500 and return that
    switch clientError.(type) {
    case errors.Http:
      break
    default:
      if c.IsAborted() {
        clientError = errors.NewHttp(c.Writer.Status(), formatErrorBody(clientError.Error()))
      } else {
        clientError = errors.NewHttp(http.StatusInternalServerError, "Unrecognized error")
      }
    }

    // Now write the error to the client
    c.JSON(clientError.(errors.Http).Code, clientError)
  }

}

func formatErrorBody(s string) string {
  return strings.Replace(s, "\n", " ", -1)
}
