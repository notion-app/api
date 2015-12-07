package suggestions

import (
	"github.com/antzucaro/matchr"
	"notion/cache"
	"notion/log"
	"notion/model"
	"time"
)

var (
	Subscriptions   = make(map[string][]*model.WsContext)
	LoopRunning     = make(map[string]bool)
	SentSuggestions = make(map[string][]model.Suggestion)
)

func Start(n model.DbNote, c *model.WsContext) {
	if _, in := Subscriptions[n.TopicId]; in {
		Subscriptions[n.TopicId] = append(Subscriptions[n.TopicId], c)
	} else {
		Subscriptions[n.TopicId] = []*model.WsContext{c}
	}
	if _, in := LoopRunning[n.TopicId]; !in {
		LoopRunning[n.TopicId] = true
		go SubLoop(n.TopicId)
	}
}

func SubLoop(topicId string) {
	timer := time.Tick(5 * time.Second)
	for range timer {
		log.Info("Finding suggestions for topic %v", topicId)
		notesInTopic := cache.GetNotesInTopic(topicId)
		for _, note := range notesInTopic {
			suggestions := FindSuggestions(note)
			for _, context := range Subscriptions[topicId] {
				for _, suggestion := range suggestions {
					SendSuggestion(suggestion, context)
				}
			}
		}
	}
}

func SendSuggestion(suggestion model.Suggestion, c *model.WsContext) {
	if suggestion.Recommendation.Genesis == c.UserId {
		log.Info("Ignoring genesis suggestion to user %v", c.UserId)
		return
	}
	if _, in := SentSuggestions[c.UserId]; in {
		for _, sentSuggestion := range SentSuggestions[c.UserId] {
			distance := matchr.Levenshtein(suggestion.Recommendation.Text, sentSuggestion.Recommendation.Text)
			if distance < 10 {
				log.Info("Ignoring edit distance violation of previous suggestion to %v", c.UserId)
				return
			}
		}
	} else {
		SentSuggestions[c.UserId] = make([]model.Suggestion, 0)
	}
	log.Info("Sending suggestion to user %v", c.UserId)
	SentSuggestions[c.UserId] = append(SentSuggestions[c.UserId], suggestion)
	SentSuggestions[suggestion.Recommendation.Genesis] = append(SentSuggestions[suggestion.Recommendation.Genesis], suggestion)
	c.SendI(suggestion)
}
