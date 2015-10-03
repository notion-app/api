// Structs which match the API responses from facebook

package model

type FacebookCurrentUser struct {
  Id string `json:"id"`
  Name string `json:"name"`
}
