package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"notion/db"
	"notion/errors"
	"notion/log"
	"notion/ws"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  2048,
		WriteBufferSize: 2048,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func EchoWebsocket(c *gin.Context) {
	conn, err := websocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if log.Error(err) {
		c.Error(errors.NewISE())
		return
	}
	for {
		typ, frame, _ := conn.ReadMessage()
		conn.WriteMessage(typ, frame)
	}
}

// Status handler for the /status route
func OpenWebsocket(c *gin.Context) {
	userId := c.MustGet("request_user_id").(string)
	in, note, err := db.GetNoteById(c.Param("note_id"))
	if log.Error(err) {
		c.Error(errors.NewISE())
		return
	}
	if !in {
		c.Error(errors.NewHttp(http.StatusNotFound, "The requested note was not found"))
		return
	}
	if userId != note.Owner {
		c.Error(errors.NewHttp(http.StatusUnauthorized, "Only owners can open websockets into their notes"))
		return
	}
	conn, err := websocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if log.Error(err) {
		c.Error(errors.NewISE())
		return
	}
	log.Info("Opening ws for user %v on %v", userId, note.Id)
	bundle := ws.NewChannelBundle()
	WrapWebsocket(conn, bundle)
	ws.ProcessMessages(bundle)
}

func WrapWebsocket(conn *websocket.Conn, bundle *ws.ChannelBundle) {
	// Read from WS, write to channel
	go func() {
		for {
			_, frameb, err := conn.ReadMessage()
			if log.Error(err) {
				bundle.Close <- true
				return
			}
			frame := make(map[string]interface{})
			err = json.Unmarshal(frameb, &frame)
			if log.Error(err) {
				continue
			}
			bundle.Incoming <- frame
		}
	}()
	// Read from channel, write to WS
	go func() {
		for {
			select {
			case msg := <-bundle.Outgoing:
				b, err := json.Marshal(msg)
				if log.Error(err) {
					continue
				}
				err = conn.WriteMessage(1, b)
				if log.Error(err) {
					continue
				}
			case <-bundle.Close:
				return
			}
		}
	}()
}
