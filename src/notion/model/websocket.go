
package model

import (
  "encoding/json"
)

func NewWSError(message string) []byte {
  j, _ := json.Marshal(map[string]interface{}{
    "type": "error",
    "message": message,
  })
  return j
}
