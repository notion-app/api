package model

import (
	"encoding/json"
)

type WsPingPong struct {
	Type string `json:"type"`
}

func NewWSError(message string) []byte {
	j, _ := json.Marshal(map[string]interface{}{
		"type":    "error",
		"message": message,
	})
	return j
}
