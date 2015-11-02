package routes

import (
  "database/sql"
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
    section, err := db.GetSectionByNotebookId(sub.NotebookId)
    if log.Error(err) {
      c.Error(errors.NewISE())
      return
    }
    course, err := db.GetCourseByCourseId(section.CourseId)
    if log.Error(err) {
      c.Error(errors.NewISE())
      return
    }
    subResponses = append(subResponses, model.NewSubscriptionResponse(sub, course, section))
  }
  c.JSON(http.StatusOK, subResponses)
}

func CreateUserSubscription(c *gin.Context) {
  userId := c.Param("user_id")
  if userId != c.MustGet("request_user_id").(string) {
    c.Error(errors.NewHttp(http.StatusUnauthorized, "Users can only create subscriptions for themselves"))
    return
  }
  var request model.SubscriptionRequest
  err := c.BindJSON(&request)
  if log.Error(err) {
    c.Error(err)
    return
  }
  sub := model.DbSubscription{
    UserId: userId,
    NotebookId: request.NotebookId,
    Name: sql.NullString{
      String: request.Name,
      Valid: true,
    },
  }
  err = db.CreateSubscription(sub)
  if log.Error(err) {
    c.Error(errors.NewISE())
    return
  }
  section, err := db.GetSectionByNotebookId(sub.NotebookId)
  if log.Error(err) {
    c.Error(errors.NewISE())
    return
  }
  course, err := db.GetCourseByCourseId(section.CourseId)
  if log.Error(err) {
    c.Error(errors.NewISE())
    return
  }
  c.JSON(http.StatusOK, model.NewSubscriptionResponse(sub, course, section))
}

func ModifyUserSubscription(c *gin.Context) {
  userId := c.Param("user_id")
  if userId != c.MustGet("request_user_id").(string) {
    c.Error(errors.NewHttp(http.StatusUnauthorized, "Users can only modify subscriptions for themselves"))
    return
  }
  var request model.SubscriptionRequest
  err := c.BindJSON(&request)
  if log.Error(err) {
    c.Error(err)
    return
  }
  sub := model.DbSubscription{
    UserId: userId,
    NotebookId: request.NotebookId,
    Name: sql.NullString{
      String: request.Name,
      Valid: true,
    },
  }
  err = db.UpdateSubscription(sub)
  if log.Error(err) {
    c.Error(errors.NewISE())
    return
  }
  section, err := db.GetSectionByNotebookId(sub.NotebookId)
  if log.Error(err) {
    c.Error(errors.NewISE())
    return
  }
  course, err := db.GetCourseByCourseId(section.CourseId)
  if log.Error(err) {
    c.Error(errors.NewISE())
    return
  }
  c.JSON(http.StatusOK, model.NewSubscriptionResponse(sub, course, section))
}

func SetUserSchool(c *gin.Context) {
  userId := c.Param("user_id")
  if userId != c.MustGet("request_user_id").(string) {
    c.Error(errors.NewHttp(http.StatusUnauthorized, "Users can only modify themselves"))
    return
  }
  var request model.ModifySchoolRequest
  err := c.BindJSON(&request)
  if log.Error(err) {
    c.Error(err)
    return
  }
  _, user, err := db.GetUserById(userId)
  if log.Error(err) {
    c.Error(err)
    return
  }
  user.School = sql.NullString{
    String: request.SchoolId,
    Valid: true,
  }
  err = db.UpdateUser(user)
  if log.Error(err) {
    c.Error(err)
    return
  }
  c.JSON(http.StatusOK, user)
}

func RemoveUserSubscription(c *gin.Context) {
  userId := c.Param("user_id")
  if userId != c.MustGet("request_user_id").(string) {
    c.Error(errors.NewHttp(http.StatusUnauthorized, "Users can only delete their own subs"))
    return
  }
  notebookId := c.Param("notebook_id")
  sub := model.DbSubscription{
    UserId: userId,
    NotebookId: notebookId,
  }
  section, err := db.GetSectionByNotebookId(sub.NotebookId)
  if log.Error(err) {
    c.Error(errors.NewISE())
    return
  }
  course, err := db.GetCourseByCourseId(section.CourseId)
  if log.Error(err) {
    c.Error(errors.NewISE())
    return
  }
  err = db.DeleteSubscription(sub)
  if log.Error(err) {
    c.Error(err)
    return
  }
  c.JSON(http.StatusOK, model.NewSubscriptionResponse(sub, course, section))
}
