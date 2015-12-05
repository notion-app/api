
package suggestions

import (
  "notion/model"
  "strings"
)

func FindSuggestions(n model.DbNote) []model.Suggestion {
  suggestions := make([]model.Suggestion, 0)
  sp := strings.Split(n.Content, "\n")
  for _, line := range sp {
    if len(line) <= 1 {
      continue
    }
    suggestions = append(suggestions, model.NewSuggestion(line, n.Owner))
  }
  return suggestions
}
