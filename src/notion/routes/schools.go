package routes

import (
	"github.com/labstack/echo"
	"net/http"
	"notion/db"
	"notion/errors"
	"notion/log"
	"notion/model"
	"notion/util"
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
	body := c.Get("BODY").(map[string]interface{})
	util.FillStruct(&request, body)
	err := validate.SchoolRequest(request)
	if log.Error(err) {
		return err
	}
	schoolRequest := model.DbSchoolRequest{
		Id:              util.NewId(),
		RequesterUserId: c.Get("TOKEN_USER_ID").(string),
		Name:            request.Name,
		Location:        request.Location,
	}
	err = db.CreateSchoolRequest(schoolRequest)
	if err != nil {
		return errors.ISE()
	}
	return nil
}

func GetCoursesForSchool(c *echo.Context) error {
	var courses []model.DbCourse
	school_id := c.Param("school_id")
	_, err := db.GenericGetMultiple("courses", "school_id", school_id, &courses)
	if log.Error(err) {
		return errors.ISE()
	}
	parsedCourses := make([]model.CourseResponse, 0)
	for _, course := range courses {
		parsedCourses = append(parsedCourses, model.CourseResponse{
			Id:     course.Id,
			Name:   course.Name,
			Number: course.Number,
		})
	}
	resp := model.CoursesForSchoolResponse{
		Courses: parsedCourses,
	}
	return c.JSON(http.StatusOK, resp)
}

func GetSectionsForCourse(c *echo.Context) error {
	var sections []model.DbCourseSection
	course_id := c.Param("course_id")
	_, err := db.GenericGetMultiple("sections", "course_id", course_id, &sections)
	if log.Error(err) {
		return errors.ISE()
	}
	resp := model.SectionsForCourseResponse{
		Sections: sections,
	}
	return c.JSON(http.StatusOK, resp)
}
