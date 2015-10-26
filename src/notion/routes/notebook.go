package routes

import (
	"github.com/labstack/echo"
	"net/http"
	"notion/logic"
)

func GetNotebookNotes(c *echo.Context) error {
	notebookId := c.Param("notebook_id")
	notes, err := logic.GetNotesByTopic(notebookId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, notes)
}
