package db

import (
	"notion/log"
	"notion/model"
)

func GetNotesInNotebook(notebookId string) ([]model.DbNote, error) {
	notes := []model.DbNote{}
	err := GenericGet(&notes, `
    select n.id, n.topic_id, n.title, n.owner, n.content, n.created_at, n.updated_at
    from notes n
    left join topics t
    on n.topic_id = t.id
    where notebook_id=$1`, notebookId)
	return notes, err
}

func GetNotesInNotebookByUser(notebookId string, userId string) ([]model.DbNote, error) {
	notes := []model.DbNote{}
	err := GenericGet(&notes, `
    select n.id, n.topic_id, n.title, n.owner, n.content, n.created_at, n.updated_at
    from notes n
    left join topics t
    on n.topic_id = t.id
    where notebook_id=$1 and owner=$2`, notebookId, userId)
	return notes, err
}

func GetUnjoinedNotesInNotebook(notebookId, userId string) ([]model.DbNote, error) {
	// Was going to try and put this logic into sql but im bad at sql
	notes, err := GetNotesInNotebook(notebookId)
	if log.Error(err) {
		return notes, err
	}
	userTopics := []string{}
	for _, note := range notes {
		if note.Owner == userId {
			userTopics = append(userTopics, note.TopicId)
		}
	}
	unjoinedNotes := []model.DbNote{}
	for _, note := range notes {
		add := true
		for _, userTopic := range userTopics {
			if note.TopicId == userTopic {
				add = false
				break
			}
		}
		if add {
			unjoinedNotes = append(unjoinedNotes, note)
		}
	}
	return unjoinedNotes, err
}

func GetNoteById(noteId string) (bool, model.DbNote, error) {
	var note model.DbNote
	in, err := GenericGetOne(&note, "select * from notes where id=$1", noteId)
	return in, note, err
}

func CreateNote(n model.DbNote) error {
	return dbmap.Insert(&n)
}

func UpdateNote(n model.DbNote) error {
	_, err := dbmap.Update(&n)
	return err
}

func DeleteNote(n model.DbNote) error {
	_, err := dbmap.Delete(&n)
	return err
}
