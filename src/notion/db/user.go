
package db

import (
  "notion/log"
  "notion/model"
)

// Gets a user by their notion-assigned Id
// Returns whether the user exists, the user model, and an error
func GetUserById(id string) (bool, model.DbUser, error) {
  var user model.DbUser
  in, useri, err := GenericGetOne("users", "id", id, &user)
  return in, useri.(model.DbUser), err
}

func GetUserByFacebookId(facebookId string) (bool, model.DbUser, error) {
  var user model.DbUser
  in, useri, err := GenericGetOne("users", "fb_user_id", facebookId, &user)
  return in, useri.(model.DbUser), err
}

func GetUserByToken(token string) (bool, model.DbUser, error) {
  var user model.DbUser
  in, useri, err := GenericGetOne("users", "fb_auth_token", token, &user)
  return in, useri.(model.DbUser), err
}

func CreateUser(u model.DbUser) error {
  log.Info("Creating new user " + u.Id)
  return dbmap.Insert(&u)
}

func UpdateUser(user model.DbUser) error {
  log.Info("Updating user token and profile pic")
  _, err := dbmap.Update(&user)
  return err
}
