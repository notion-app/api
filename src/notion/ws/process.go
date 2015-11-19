package ws

import (
	"fmt"
	"notion/log"
	"notion/model"
)

func ProcessMessages(incoming chan map[string]interface{}, outgoing chan map[string]interface{}) {
	for msg := range incoming {
		err := DispatchFrame(msg, outgoing)
		if log.Error(err) {
			continue
		}
	}
}

func DispatchFrame(frame map[string]interface{}, outgoing chan map[string]interface{}) error {
	fType, in := frame["type"]
  if !in {
    return fmt.Errorf("Must provide type tag in websocket body")
  }
  var fTypeS string
  switch fType.(type) {
  case string:
    fTypeS = fType.(string)
  default:
    return fmt.Errorf("WS message type must be a string")
  }
  switch fTypeS {
  case "ping":
		HandlePing(model.WsPingPong{
      Type: "ping",
    }, outgoing)
  case "pong":
		HandlePing(model.WsPingPong{
      Type: "pong",
    }, outgoing)
	default:
		return fmt.Errorf("Unrecognized type passed through websocket; doing nothing")
  }
  return nil
}

func HandlePing(p model.WsPingPong, outgoing chan map[string]interface{}) {
	outgoing <- map[string]interface{}{
		"type": "pong",
	}
}
