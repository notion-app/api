package routes

import (
	"database/sql"
	"github.com/labstack/echo"
	"net/http"
	"notion/db"
	"notion/errors"
	"notion/log"
	"notion/model"
	"notion/util"
	"notion/validate"
)

func GetUser(c *echo.Context) error {
	userId := c.Param("user_id")
	if userId != c.Get("TOKEN_USER_ID") {
		return errors.Unauthorized("notion")
	}
	in, user, err := db.GetUserById(userId)
	if err != nil {
		return errors.ISE()
	}
	if !in {
		return errors.NotFound()
	}
	return c.JSON(http.StatusOK, user)
}

func GetUsersSubscriptions(c *echo.Context) error {
	userId := c.Param("user_id")
	if userId != c.Get("TOKEN_USER_ID") {
		return errors.Unauthorized("notion")
	}
	subscriptions, err := db.GetSubscriptionsByUserId(userId)
	if err != nil {
		return errors.ISE()
	}
	return c.JSON(http.StatusOK, model.UserSubscriptionsResponse{
		Subscriptions: subscriptions,
	})
}

func CreateUserSubscription(c *echo.Context) error {
	var request model.SubscriptionRequest
	userId := c.Param("user_id")
	if userId != c.Get("TOKEN_USER_ID") {
		return errors.Unauthorized("notion")
	}
	body := c.Get("BODY").(map[string]interface{})
	util.FillStruct(&request, body)
	err := validate.AddSubscriptionRequest(request)
	if log.Error(err) {
		return err
	}
	sub := model.DbSubscription{
		UserId:     userId,
		NotebookId: request.NotebookId,
	}
	err = db.CreateSubscription(sub)
	if log.Error(err) {
		return errors.ISE()
	}
	return nil
}

func RemoveUserSubscription(c *echo.Context) error {
	var request model.SubscriptionRequest
	userId := c.Param("user_id")
	if userId != c.Get("TOKEN_USER_ID") {
		return errors.Unauthorized("notion")
	}
	body := c.Get("BODY").(map[string]interface{})
	util.FillStruct(&request, body)
	err := validate.RemoveSubscriptionRequest(request)
	if log.Error(err) {
		return err
	}
	sub := model.DbSubscription{
		UserId:     userId,
		NotebookId: request.NotebookId,
	}
	err = db.RemoveSubscription(sub)
	if log.Error(err) {
		return errors.ISE()
	}
	return nil
}

func SetUserSchool(c *echo.Context) error {
	var request model.AddSchoolRequest
	userId := c.Param("user_id")
	if userId != c.Get("TOKEN_USER_ID") {
		return errors.Unauthorized("notion")
	}
	body := c.Get("BODY").(map[string]interface{})
	util.FillStruct(&request, body)

	_, user, err := db.GetUserById(userId)
	if log.Error(err) {
		return err
	}

	user.School = sql.NullString{
		String: request.SchoolId,
		Valid:  true,
	}
	err = db.UpdateUser(user)
	if log.Error(err) {
		return err
	}

	return nil

}
