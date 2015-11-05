package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"notion/log"
	"notion/model"
)

func Loop(conn *websocket.Conn, typ int, frameb []byte, err error) {
	frame := make(map[string]interface{})
	err1 := json.Unmarshal(frameb, &frame)
	if log.Error(err1) {
		conn.WriteMessage(1, model.NewWSError("Message provided is not valid json"))
		return
	}

}
