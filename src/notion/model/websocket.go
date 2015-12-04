package model

import (
	"encoding/json"
	"notion/ot"
)

type WsContext struct {
	UserId   string
	NoteId   string
	Incoming chan map[string]interface{}
	Outgoing chan map[string]interface{}
	Close    chan bool
}

func NewContext(userId string, noteId string) *WsContext {
	return &WsContext{
		UserId:   userId,
		NoteId:   noteId,
		Incoming: make(chan map[string]interface{}),
		Outgoing: make(chan map[string]interface{}),
		Close:    make(chan bool),
	}
}

func (cb *WsContext) SendError(message string) {
	cb.SendM(map[string]interface{}{
		"type":    "error",
		"message": message,
	})
}

func (cb *WsContext) SendM(m map[string]interface{}) {
	cb.Outgoing <- m
}

type WsPingPong struct {
	Type string `json:"type"`
}

type WsUpdate struct {
	Type   string `json:"type"`
	Update struct {
		BaseLength   int          `json:"baseLength"`
		TargetLength int          `json:"targetLength"`
		Operations   ot.Transform `json:"ops"`
	} `json:"update"`
}

func NewWSError(message string) []byte {
	j, _ := json.Marshal(map[string]interface{}{
		"type":    "error",
		"message": message,
	})
	return j
}
