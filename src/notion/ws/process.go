package ws

import (
	"fmt"
	"notion/db"
	"notion/log"
	"notion/model"
)

var (
	SubscriptionMap = make(map[string][]*model.WsContext)
	NoteContent     = make(map[string]model.DbNote)
)

func ProcessMessages(bundle *model.WsContext) {

	// Load in the initial copy of the content of the note into memory
	// Theoretically this would just balloon in memory usage because we never
	// clear these out when a WsContext is closed, but isn't that what heroku's
	// auto-restarting thing is for anyway? No?
	if _, in := NoteContent[bundle.NoteId]; !in {
		in, note, err := db.GetNoteById(bundle.NoteId)
		if log.Error(err) || !in {
			return
		}
		NoteContent[bundle.NoteId] = note
	}

	// Cache the subscription context so we can send and receive updates
	if _, in := SubscriptionMap[bundle.NoteId]; in {
		SubscriptionMap[bundle.NoteId] = append(SubscriptionMap[bundle.NoteId], bundle)
	} else {
		SubscriptionMap[bundle.NoteId] = []*model.WsContext{bundle}
	}

	// Start iterating over each incoming websocket message
	for msg := range bundle.Incoming {
		err := DispatchFrame(msg, bundle)
		if log.Error(err) {
			bundle.SendError(err.Error())
			continue
		}
	}
}

func DispatchFrame(frame map[string]interface{}, bundle *model.WsContext) error {
	fType, in := frame["type"]
	if !in {
		return fmt.Errorf("Must provide type tag in websocket body")
	}
	var fTypeS string
	switch fType.(type) {
	case string:
		fTypeS = fType.(string)
	default:
		return fmt.Errorf("Message type provided must be a string")
	}
	switch fTypeS {
	case "ping":
		return HandlePing(model.WsPingPong{
			Type: "ping",
		}, bundle)
	case "pong":
		return HandlePing(model.WsPingPong{
			Type: "pong",
		}, bundle)
	case "update":
		return HandleUpdate(frame, bundle)
	default:
		return fmt.Errorf("Unrecognized message type; doing nothing")
	}
	return nil
}

func HandlePing(p model.WsPingPong, bundle *model.WsContext) error {
	bundle.Outgoing <- map[string]interface{}{
		"type": "pong",
	}
	return nil
}
