package ws

import (
	"notion/cache"
	"notion/log"
	"notion/model"
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
	note := cache.GetNote(c.NoteId)
	note.Content, err = updateFrame.Update.Operations.Apply(note.Content)
	if log.Error(err) {
		return err
	}
	cache.Note(note)

	return nil

}
