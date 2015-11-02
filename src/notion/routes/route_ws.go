package routes

import (
	"github.com/gin-gonic/gin"
  "github.com/gorilla/websocket"
  "net/http"
  "notion/db"
  "notion/errors"
  "notion/log"
  // "notion/ws"
)

var (
	websocketUpgrader = websocket.Upgrader{
    ReadBufferSize: 2048,
    WriteBufferSize: 2048,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
  }
)

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
	ch := make(chan map[string]interface{})
	WrapWebsocket(conn, ch)
  // for {
	// 	typ, frame, err := conn.ReadMessage()
	// 	go ws.Loop(conn, typ, frame, err)
  // }
}

func WrapWebsocket(conn *websocket.Conn, ch chan map[string]interface{}) {
	// reader
	go func() {
		for {
			conn.ReadMessage()
		}
	}()
	// writer
	go func() {
		select {

		}
	}()
}
