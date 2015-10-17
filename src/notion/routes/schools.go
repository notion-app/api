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
