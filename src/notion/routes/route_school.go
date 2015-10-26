
package routes

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "notion/db"
  "notion/errors"
)

func GetAllSchools(c *gin.Context) {
  schools, err := db.GetAllSchools()
  if err != nil {
    c.Error(errors.NewISE())
    return
  }
  c.JSON(http.StatusOK, schools)
}
