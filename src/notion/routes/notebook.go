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
	userFilter := c.Query("user")
	if userFilter != "" {
		for _, topic := range notes {
			remove := []int{}
			for note_i, _ := range topic.Notes {
				if topic.Notes[note_i].Owner != userFilter {
					remove = append(remove, note_i)
				}
			}
			for _, i := range remove {
				topic.Notes = append(topic.Notes[:i], topic.Notes[i+1:]...)
			}
		}
	}
	// unjoinedFilter := c.Query("unjoined")
	return c.JSON(http.StatusOK, notes)
}
