
package routes

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "notion/db"
  "notion/errors"
  "notion/log"
  "notion/model"
)

func GetAllSchools(c *gin.Context) {
  schools, err := db.GetAllSchools()
  if log.Error(err) {
    c.Error(errors.NewISE())
    return
  }
  c.JSON(http.StatusOK, schools)
}

func GetCoursesForSchool(c *gin.Context) {
  schoolId := c.Param("school_id")
  courses, err := db.GetCoursesForSchool(schoolId)
  if log.Error(err) {
    c.Error(errors.NewISE())
    return
  }
  courseResponses := make([]model.CourseResponse, 0)
  for _, course := range courses {
    courseResponses = append(courseResponses, model.CourseResponseWithoutSchool(course))
  }
  c.JSON(http.StatusOK, courseResponses)
}
