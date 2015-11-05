package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notion/db"
	"notion/errors"
	"notion/log"
	"notion/model"
	"notion/util"
)

func GetSingleSchool(c *gin.Context) {
	schoolId := c.Param("school_id")
	in, school, err := db.GetSchool(schoolId)
	if log.Error(err) {
		c.Error(errors.NewISE())
		return
	}
	if !in {
		c.Error(errors.NewHttp(http.StatusNotFound, "School requested could not be found"))
		return
	}
	c.JSON(http.StatusOK, school)
}

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

func GetSectionsForCourse(c *gin.Context) {
	courseId := c.Param("course_id")
	sections, err := db.GetSectionsForCourse(courseId)
	if log.Error(err) {
		c.Error(errors.NewISE())
		return
	}
	sectionsResponse := make([]model.SectionResponse, 0)
	for _, section := range sections {
		sectionsResponse = append(sectionsResponse, model.SectionResponseWithoutCourse(section))
	}
	c.JSON(http.StatusOK, sectionsResponse)
}

func PostSchoolRequest(c *gin.Context) {
	var request model.SchoolRequestRequest
	err := c.BindJSON(&request)
	if log.Error(err) {
		c.Error(err)
		return
	}
	dbreq := model.DbSchoolRequest{
		Id:              util.NewId(),
		RequesterUserId: c.MustGet("request_user_id").(string),
		Name:            request.Name,
		Location:        request.Location,
	}
	err = db.CreateSchoolRequest(dbreq)
	if log.Error(err) {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dbreq)
}
