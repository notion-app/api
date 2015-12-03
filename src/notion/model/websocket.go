package model

import (
	"encoding/json"
	"notion/ot"
)

type WsPingPong struct {
	Type string `json:"type"`
}

type WsUpdate struct {
	Type string `json:"type"`
	Update struct {
		BaseLength int `json:"baseLength"`
		TargetLength int `json:"targetLength"`
		Operations ot.Transform `json:"ops"`
	} `json:"update"`
}

func NewWSError(message string) []byte {
	j, _ := json.Marshal(map[string]interface{}{
		"type":    "error",
		"message": message,
	})
	return j
}
