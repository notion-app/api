package cache

import (
	"notion/db"
	"notion/log"
	"notion/model"
	"time"
)

var (
	NoteIdCache    = make(map[string]*model.DbNote)
	NoteTopicCache = make(map[string]map[string]*model.DbNote)
)

func Note(n model.DbNote) {
	if oldNote, in := NoteIdCache[n.Id]; in {
		log.Error(db.UpdateNote(*oldNote))
	}
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

func GetNotesInTopic(topicId string) []model.DbNote {
	l := make([]model.DbNote, 0)
	for _, note := range NoteTopicCache[topicId] {
		l = append(l, *note)
	}
	return l
}

func DumpNoteCache() {
	ticker := time.Tick(5 * time.Second)
	for range ticker {
		for _, note := range NoteIdCache {
			if log.Error(db.UpdateNote(*note)) {
				continue
			}
		}
	}
}
