package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"notion/db"
	"notion/errors"
	"notion/log"
	"notion/model"
	"notion/ot"
	"notion/util"
	"sync"
	"time"
)

var (
	// lol i cant fully comprehend how dumb this right here is
	NoteLockHash = make(map[string]*sync.Mutex)
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
	err := c.BindJSON(&request)
	if log.Error(err) {
		c.Error(err)
		return
	}
	userId := c.MustGet("request_user_id").(string)
	notebookId := c.Param("notebook_id")
	now := time.Now()
	dbn := model.DbNote{
		Id:        util.NewId(),
		Owner:     userId,
		Content:   "",
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	if request.TopicId == "" {
		dbn.TopicId = util.NewId()
		err = db.CreateTopic(model.DbTopic{
			Id:         dbn.TopicId,
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
			Valid:  true,
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
	c.JSON(http.StatusOK, model.TopicResponse{
		Id: dbn.TopicId,
		Notes: []model.FullNoteResponse{
			model.NewFullNoteResponse(dbn),
		},
	})
}

func ModifyNote(c *gin.Context) {
	var request model.ModifyNoteRequest
	err := c.BindJSON(&request)
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
	now := time.Now()
	dbn.UpdatedAt = &now

	if mu, in := NoteLockHash[dbn.Id]; in {
		mu.Lock()
	} else {
		NoteLockHash[dbn.Id] = &sync.Mutex{}
		NoteLockHash[dbn.Id].Lock()
	}
	err = db.UpdateNote(dbn)
	NoteLockHash[dbn.Id].Unlock()

	if log.Error(err) {
		c.Error(errors.NewISE())
		return
	}
	c.JSON(http.StatusOK, model.TopicResponse{
		Id: dbn.TopicId,
		Notes: []model.FullNoteResponse{
			model.NewFullNoteResponse(dbn),
		},
	})
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
	c.JSON(http.StatusOK, model.TopicResponse{
		Id: dbn.TopicId,
		Notes: []model.FullNoteResponse{
			model.NewFullNoteResponse(dbn),
		},
	})
}

func PostNoteChange(c *gin.Context) {
	var ott ot.Transform
	noteId := c.Param("note_id")
	userId := c.MustGet("request_user_id")
	err := c.BindJSON(&ott)
	if log.Error(err) {
		c.Error(err)
		return
	}

	in, dbn, err := db.GetNoteById(noteId)
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

	if mu, in := NoteLockHash[dbn.Id]; in {
		mu.Lock()
	} else {
		NoteLockHash[dbn.Id] = &sync.Mutex{}
		NoteLockHash[dbn.Id].Lock()
	}
	dbn.Content, err = ott.Do(dbn.Content)
	if log.Error(err) {
		c.Error(errors.NewHttp(http.StatusBadRequest, err.Error()))
		return
	}
	err = db.UpdateNote(dbn)
	NoteLockHash[dbn.Id].Unlock()
	if log.Error(err) {
		c.Error(errors.NewISE())
		return
	}

	c.JSON(http.StatusOK, dbn)

}
