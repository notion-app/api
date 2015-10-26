
package routes

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "notion/db"
  "notion/errors"
  "notion/log"
  "notion/model"
)

func GetNotebookNotes(c *gin.Context) {
  userId := c.MustGet("request_user_id").(string)
  notebookId := c.Param("notebook_id")

  // Check query params
  filterUserId := c.Query("user")
  filterUnjoined := c.Query("unjoined")

  // And execute those query params
  var notes []model.DbNote
  var err error
  if filterUserId == "" && filterUnjoined == "" {
    notes, err = db.GetNotesInNotebook(notebookId)
  } else if filterUserId == "" {
    notes, err = db.GetUnjoinedNotesInNotebook(notebookId, userId)
  } else if filterUnjoined == "" {
    notes, err = db.GetNotesInNotebookByUser(notebookId, userId)
  } else {
    c.Error(errors.NewHttp(http.StatusBadRequest, "Cannot provide both unjoined and user parameters lol"))
    return
  }

  if log.Error(err) {
    c.Error(errors.NewISE())
    return
  }

  // Transform the resultant list into hierarchial [topic -> []note]
  topicHash := make(map[string][]model.NoteResponse)
  for _, dbnote := range notes {
    if ar, in := topicHash[dbnote.TopicId]; in {
      topicHash[dbnote.TopicId] = append(ar, model.NewNoteResponse(dbnote))
    } else {
      topicHash[dbnote.TopicId] = []model.NoteResponse{
        model.NewNoteResponse(dbnote),
      }
    }
  }
  topicResponses := []model.TopicResponse{}
  for topicid, notelist := range topicHash {
    topicResponses = append(topicResponses, model.TopicResponse{
      Id: topicid,
      Notes: notelist,
    })
  }
  c.JSON(http.StatusOK, topicResponses)
}
