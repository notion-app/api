
package ws

type ChannelBundle struct {
  Incoming chan map[string]interface{}
  Outgoing chan map[string]interface{}
  Close chan bool
}

func NewChannelBundle() *ChannelBundle {
  return &ChannelBundle{
    Incoming: make(chan map[string]interface{}),
    Outgoing: make(chan map[string]interface{}),
    Close: make(chan bool),
  }
}

func (cb *ChannelBundle) Send(m map[string]interface{}) {
  cb.Outgoing <- m
}
