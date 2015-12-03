
package ws

import (
  // "fmt"
  // "notion/log"
  // "notion/model"
)

func HandleUpdate(frame map[string]interface{}, c *Context) {
  // _, err := ParseUpdateFrame(frame)
  // if log.Error(err) {
  //   c.SendError(err.Error())
  //   return
  // }
}

// func ParseUpdateFrame(frame map[string]interface{}) (model.WsUpdate, error) {
//   var up model.WsUpdate
//   updateObj, in := frame["update"]
//   if !in {
//     return up, fmt.Errorf("Must provide 'update' field in websocket request")
//   }
//   baseLength, in := updateObj.(map[string]interface{})[""]
//
// }
