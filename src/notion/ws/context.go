
package ws

type Context struct {
  UserId string
  NoteId string
  Incoming chan map[string]interface{}
  Outgoing chan map[string]interface{}
  Close chan bool
}

func NewContext(userId string, noteId string) *Context {
  return &Context{
    UserId: userId,
    NoteId: noteId,
    Incoming: make(chan map[string]interface{}),
    Outgoing: make(chan map[string]interface{}),
    Close: make(chan bool),
  }
}

func (cb *Context) Send(m map[string]interface{}) {
  cb.Outgoing <- m
}
