package db

import (
	"database/sql"
	"fmt"
	"notion/log"
	"notion/model"
)

func GetNotes(notebookId string) ([]model.DbNote, error) {
  var notes []model.DbNote
	query := fmt.Sprintf(`
		select n.id, n.topic_id, n.title, n.owner, n.content, n.created_at, n.updated_at
		from notes n
		left join topics t
		on n.topic_id = t.id
		where notebook_id = '%v'
	`, notebookId)
	_, err := dbmap.Select(&notes, query)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return notes, nil
		default:
			log.Error(err)
			return nil, err
		}
	}
  return notes, nil
}
