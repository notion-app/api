
package model

import (
  "notion/util"
)

type Suggestion struct {
  Type string `json:"type"`
  Recommendation InnerSuggestion `json:"recommendation"`
}

type InnerSuggestion struct {
  Text string `json:"text"`
  Genesis string `json:"genesis"`
  Id string `json:"id"`
}

func NewSuggestion(text string, from string) Suggestion {
  return Suggestion{
    Type: "recommendation",
    Recommendation: InnerSuggestion{
      Id: util.NewId(),
      Text: text,
      Genesis: from,
    },
  }
}
