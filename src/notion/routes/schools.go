package routes

import (
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"notion/db"
	"notion/errors"
	"notion/log"
	"notion/logic"
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
	user, err := logic.AuthenticateNotionUser(c.Param("token"))
	if log.Error(err) {
		return err
	}
	schoolRequest := model.DbSchoolRequest{
		Id:              util.NewId(),
		RequesterUserId: user.Id,
		Name:            request.Name,
		Location:        request.Location,
	}
	err = db.CreateSchoolRequest(schoolRequest)
	if err != nil {
		return errors.ISE()
	}
	return nil
}
