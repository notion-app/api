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
  for {
    // ws.Loop(conn)
    t, msg, err := conn.ReadMessage()
    if log.Error(err) {
      break
    }
    conn.WriteMessage(t, msg)
  }
}
