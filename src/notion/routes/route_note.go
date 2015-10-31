package routes

import (
  "database/sql"
	"github.com/gin-gonic/gin"
  "net/http"
  "notion/db"
  "notion/errors"
  "notion/log"
  "notion/model"
  "notion/util"
  "time"
)

func GetSingleNote(c *gin.Context) {
  noteId := c.Param("note_id")
  in, dbn, err := db.GetNoteById(noteId)
  if err != nil {
    c.Error(errors.NewISE())
    return
  }
  if !in {
    c.Error(errors.NewHttp(http.StatusNotFound, "The requested note could not be found"))
    return
  }
  c.JSON(http.StatusOK, model.NewFullNoteResponse(dbn))
}

func CreateNote(c *gin.Context) {
  var request model.CreateNoteRequest
  err := c.Bind(&request)
  if log.Error(err) {
    c.Error(err)
    return
  }
  userId := c.MustGet("request_user_id").(string)
  notebookId := c.Param("notebook_id")
  now := time.Now()
  dbn := model.DbNote{
    Id: util.NewId(),
    Owner: userId,
    Content: "",
    CreatedAt: &now,
    UpdatedAt: &now,
  }
  if request.TopicId == "" {
    dbn.TopicId = util.NewId()
    err = db.CreateTopic(model.DbTopic{
      Id: dbn.TopicId,
      NotebookId: notebookId,
    })
  } else {
    dbn.TopicId = request.TopicId
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
  if log.Error(err) {
    c.Error(err)
    return
  }
  err = db.CreateNote(dbn)
  if log.Error(err) {
    c.Error(err)
    return
  }
	c.JSON(http.StatusOK, model.NewFullNoteResponse(dbn))
}

func ModifyNote(c *gin.Context) {
  var request model.ModifyNoteRequest
  err := c.Bind(&request)
  if err != nil {
    c.Error(err)
    return
  }
  userId := c.MustGet("request_user_id")
  notebookId := c.Param("note_id")
  in, dbn, err := db.GetNoteById(notebookId)
  if err != nil {
    c.Error(errors.NewISE())
    return
  }
  if !in {
    c.Error(errors.NewHttp(http.StatusNotFound, "The requested note could not be found"))
    return
  }
  if userId != dbn.Owner {
    c.Error(errors.NewHttp(http.StatusUnauthorized, "Only the owner of a note can modify its content"))
    return
  }
  if request.TopicId != "" {
    dbn.TopicId = request.TopicId
  }
  if request.Title != "" {
    dbn.Title.Valid = true
    dbn.Title.String = request.Title
  }
  if request.Content != "" {
    dbn.Content = request.Content
  }
  err = db.UpdateNote(dbn)
  if log.Error(err) {
    c.Error(errors.NewISE())
    return
  }
  c.JSON(http.StatusOK, model.NewFullNoteResponse(dbn))
}

func DeleteNote(c *gin.Context) {
  userId := c.MustGet("request_user_id")
  notebookId := c.Param("note_id")
  in, dbn, err := db.GetNoteById(notebookId)
  if err != nil {
    c.Error(errors.NewISE())
    return
  }
  if !in {
    c.Error(errors.NewHttp(http.StatusNotFound, "The requested note could not be found"))
    return
  }
  if userId != dbn.Owner {
    c.Error(errors.NewHttp(http.StatusUnauthorized, "Only the owner of a note can modify its content"))
    return
  }
  err = db.DeleteNote(dbn)
  if log.Error(err) {
    c.Error(errors.NewISE())
    return
  }
  c.JSON(http.StatusOK, model.NewFullNoteResponse(dbn))
}
