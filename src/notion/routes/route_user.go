package routes

import (
  "github.com/gin-gonic/gin"
	"net/http"
  "notion/db"
  "notion/errors"
  "notion/log"
  "notion/model"
)

func GetUser(c *gin.Context) {
  userId := c.Param("user_id")
  if userId != c.MustGet("request_user_id").(string) {
    c.Error(errors.NewHttp(http.StatusUnauthorized, "Users can only view detailed information about themselves"))
    return
  }
  _, user, err := db.GetUserById(userId)
  if log.Error(err) {
    c.Error(errors.NewISE())
    return
  }
  c.JSON(http.StatusOK, model.NewUserResponse(user))
}

func GetUsersSubscriptions(c *gin.Context) {
  userId := c.Param("user_id")
  if userId != c.MustGet("request_user_id").(string) {
    c.Error(errors.NewHttp(http.StatusUnauthorized, "Users can only view the subscriptions of themselves"))
    return
  }
  subs, err := db.GetUserSubscriptions(userId)
  if log.Error(err) {
    c.Error(errors.NewISE())
    return
  }
  subResponses := make([]model.SubscriptionResponse, 0)
  for _, sub := range subs {
    subResponses = append(subResponses, model.NewSubscriptionResponse(sub))
  }
  c.JSON(http.StatusOK, subResponses)
}
