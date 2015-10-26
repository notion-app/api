
package errors

import (
  "fmt"
  "net/http"
)

const (
  ISE = http.StatusInternalServerError
)

type Http struct {
  Code int `json:"code"`
  Message string `json:"message"`
}

func (h Http) Error() string {
  return fmt.Sprintf("Error (%v) : %v", h.Code, h.Message)
}

func NewHttp(code int, message string) Http {
  return Http{
    Code: code,
    Message: message,
  }
}

func NewISE() Http {
  return Http{
    Code: http.StatusInternalServerError,
    Message: "Internal Server Error",
  }
}
