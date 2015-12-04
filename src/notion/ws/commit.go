
package ws

import (
  "notion/db"
  "notion/log"
  "time"
)

// Lol this is terrible
func InitCommitter() {
  log.Info("Starting in-memory cache database commit routine")
  go func() {
    ticker := time.Tick(10 * time.Second)
    for range ticker {
      if len(NoteContent) > 0 {
        log.Info("Dumping in-memory note-cache to database")
      }
      for _, content := range NoteContent {
        if log.Error(db.UpdateNote(content)) {
          continue
        }
      }
    }
  }()
}
