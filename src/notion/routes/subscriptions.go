package routes

import (
	"github.com/labstack/echo"
	"net/http"
	"notion/db"
	"notion/errors"
	"notion/model"
)

func GetSubscriptions(c *echo.Context) error {
	userId := c.Query("user_id")
	subscriptions, err := db.GetSubscriptionsByUserId(userId)
	if err != nil {
		return errors.ISE()
	}
	return c.JSON(http.StatusOK, model.UserSubscriptionsResponse{
		Subscriptions: subscriptions,
	})
}
