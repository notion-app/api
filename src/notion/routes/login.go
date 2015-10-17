package routes

import (
	"github.com/labstack/echo"
	"notion/log"
	"notion/logic"
	"notion/model"
	"notion/util"
	"notion/validate"
)

func Login(c *echo.Context) error {
	var request model.LoginRequest
	body := c.Get("BODY").(map[string]interface{})
	util.FillStruct(&request, body)
	err := validate.Login(request)
	if log.Error(err) {
		return err
	}
	code, resp, err := logic.DoUserCreateOrLogin(request)
	if log.Error(err) {
		return err
	}
	return c.JSON(code, resp)
}
