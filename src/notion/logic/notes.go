package logic

import (
	"fmt"
	"notion/db"
	"notion/errors"
	"notion/model"
)

func GetNotesByTopic(notebookId string) ([]model.TopicResponse, error) {
	notes, err := db.GetNotes(notebookId)
	if err != nil {
		return nil, errors.ISE()
	}
	topicHash := make(map[string][]model.NoteResponse)
	for _, note := range notes {
		if ar, in := topicHash[note.TopicId]; in {
			topicHash[note.TopicId] = append(ar, model.NewNoteResponse(note))
		} else {
			response := model.NewNoteResponse(note)
			topicHash[note.TopicId] = []model.NoteResponse{}
			topicHash[note.TopicId] = append(topicHash[note.TopicId], response)
		}
	}
	topicResponses := []model.TopicResponse{}
	for topicId, noteList := range topicHash {
		tr := model.TopicResponse{
			Id:    topicId,
			Notes: noteList,
		}
		topicResponses = append(topicResponses, tr)
	}
	return topicResponses, nil
}
