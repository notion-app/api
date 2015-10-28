package routes

import (
  "database/sql"
	"github.com/gin-gonic/gin"
  "net/http"
  "notion/db"
  "notion/log"
  "notion/model"
  "notion/util"
  "time"
)

func CreateNote(c *gin.Context) {
  var request model.NoteRequest
  err := c.Bind(&request)
  if log.Error(err) {
    c.Error(err)
    return
  }
  userId := c.MustGet("request_user_id").(string)
  now := time.Now()
  dbn := model.DbNote{
    Id: util.NewId(),
    Owner: userId,
    Content: "",
    CreatedAt: &now,
    UpdatedAt: &now,
  }
  if request.TopicId == "" {
    dbn.TopicId = sql.NullString{
      Valid: false,
    }
  } else {
    dbn.TopicId = sql.NullString{
      Valid: true,
      String: request.TopicId,
    }
  }
  if request.Title == "" {
    dbn.Title = sql.NullString{
      Valid: false,
    }
  } else {
    dbn.Title = sql.NullString{
      Valid: true,
      String: request.Title,
    }
  }
  err = db.CreateNote(dbn)
	c.JSON(http.StatusOK, model.NewFullNoteResponse(dbn))
}
