package routes

import (
	"github.com/labstack/echo"
	"net/http"
	"notion/db"
	"notion/errors"
	"notion/model"
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
