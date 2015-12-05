
package cache

import (
  "notion/db"
  "notion/log"
  "notion/model"
  "time"
)

var (
  NoteIdCache = make(map[string]*model.DbNote)
  NoteTopicCache = make(map[string]map[string]*model.DbNote)
)

func Note(n model.DbNote) {
  NoteIdCache[n.Id] = &n
  if _, in := NoteTopicCache[n.TopicId]; in {
    NoteTopicCache[n.TopicId][n.Id] = &n
  } else {
    NoteTopicCache[n.TopicId] = make(map[string]*model.DbNote)
    NoteTopicCache[n.TopicId][n.Id] = &n
  }
}

func GetNote(id string) model.DbNote {
  return *NoteIdCache[id]
}

func DumpNoteCache() {
  ticker := time.Tick(10 * time.Second)
  for range ticker {
    for _, note := range NoteIdCache {
      if log.Error(db.UpdateNote(*note)) {
        continue
      }
    }
  }
}
