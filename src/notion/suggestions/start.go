
package suggestions

import (
  "github.com/antzucaro/matchr"
  "notion/cache"
  "notion/model"
)

var (
  Subscriptions = make(map[string][]*model.WsContext)
  LoopRunning = make(map[string]bool)
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
  for {
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
    return
  }
  if _, in := SentSuggestions[c.UserId]; in {
    for _, sentSuggestion := range SentSuggestions[c.UserId] {
      distance := matchr.Levenshtein(suggestion.Recommendation.Text, sentSuggestion.Recommendation.Text)
      if distance < 10 {
        return
      }
    }
  } else {
    SentSuggestions[c.UserId] = make([]model.Suggestion, 0)
  }
  SentSuggestions[c.UserId] = append(SentSuggestions[c.UserId], suggestion)
  c.SendI(suggestion)
}