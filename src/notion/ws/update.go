package ws

import (
	"fmt"
	"notion/log"
	"notion/model"
	"notion/recs"
	"notion/util"
)

func HandleUpdate(frame map[string]interface{}, c *model.WsContext) error {
	var updateFrame model.WsUpdate
	err := util.FillStruct(&updateFrame, frame)
	if log.Error(err) {
		return err
	}
	log.Info("Processing update %+v", updateFrame)

	// Update the note content in our cache
	if note, in := NoteContent[c.NoteId]; in {
		note.Content, err = updateFrame.Update.Operations.Apply(note.Content)
		if err != nil {
			return err
		}
		NoteContent[c.NoteId] = note
	} else {
		// If its not in our cache then shit is fucked
		log.Error(fmt.Errorf("Note is not cached locally, this shouldn't happen"))
	}

	// Register only the text additions in the recommendation engine
	textAt := 0
	for _, op := range updateFrame.Update.Operations {
		switch op.(type) {
		case float64:
			if int(op.(float64)) >= 0 {
				textAt = int(op.(float64))
			}
		case string:
			go recs.Register(op.(string), textAt, c.UserId, c.NoteId, c.Outgoing)
			textAt += len(op.(string))
		}
	}

	return nil
}
