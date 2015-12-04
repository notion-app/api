
package ws

import (
  "fmt"
  "notion/log"
  "notion/model"
  "notion/util"
)

func HandleUpdate(frame map[string]interface{}, c *Context) error {
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

  return nil
}
