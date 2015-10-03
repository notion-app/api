
package service

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "notion/log"
  "notion/model"
)

type Facebook struct {}

// Returns information about the currently logged in user
//  bool: whether or not the authToken is valid
//  struct: the current user information
//  error: these would always represent ISE's
func (f Facebook) GetCurrentUser(authToken string) (bool, model.FacebookCurrentUser, error) {
  var data model.FacebookCurrentUser
  resp, err := http.Get(fmt.Sprintf("https://graph.facebook.com/v2.4/me?access_token=%v", authToken))
  if log.Error(err) {
    return false, data, err
  }
  if resp.StatusCode == 400 {
    return false, data, nil
  }
  bytes, err := ioutil.ReadAll(resp.Body)
  if log.Error(err) {
    return false, data, err
  }
  err = json.Unmarshal(bytes, &data)
  if log.Error(err) {
    return false, data, err
  }
  return true, data, nil
}

// Calls a facebook api endpoint to make a short-lived token long-lived
// Returns a boolean as to whether or not the token provided is valid
func (f Facebook) ExtendToken(authToken string) (bool, error) {
  return false, nil
}
