package ws

import (
	"fmt"
	"notion/log"
	"notion/model"
)

func ProcessMessages(bundle *ChannelBundle) {
	for msg := range bundle.Incoming {
		err := DispatchFrame(msg, bundle)
		if log.Error(err) {
			continue
		}
	}
}

func DispatchFrame(frame map[string]interface{}, bundle *ChannelBundle) error {
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
    }, bundle)
  case "pong":
		HandlePing(model.WsPingPong{
      Type: "pong",
    }, bundle)
	default:
		return fmt.Errorf("Unrecognized type passed through websocket; doing nothing")
  }
  return nil
}

func HandlePing(p model.WsPingPong, bundle *ChannelBundle) {
	bundle.Outgoing <- map[string]interface{}{
		"type": "pong",
	}
}
