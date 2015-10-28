
package db

import (
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
  notes := []model.DbNote{}
  err := GenericGet(&notes, `
    select n.id, n.topic_id, n.title, n.owner, n.content, n.created_at, n.updated_at
    from notes n
    left join topics t
    on n.topic_id = t.id
    where notebook_id = $1
    except
    select n2.id, n2.topic_id, n2.title, n2.owner, n2.content, n2.created_at, n2.updated_at
    from notes n2
    left join topics t2
    on n2.topic_id = t2.id
    where owner = $2`, notebookId, userId)
  return notes, err
}

func CreateNote(n model.DbNote) error {
  return dbmap.Insert(&n)
}
