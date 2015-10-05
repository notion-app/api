
package db

import (
  "database/sql"
  "notion/log"
  "notion/model"
)

// Gets a user by their notion-assigned Id
// Returns whether the user exists, the user model, and an error
func GetUserById(id string) (bool, model.DbUser, error) {
  log.Info("Getting user " + id)
  var user model.DbUser
  err := dbmap.SelectOne(&user, "select * from users where id=?", id)
  if err != nil {
    switch err {
    case sql.ErrNoRows:
      return false, user, nil
    default:
      return false, user, err
    }
  }
  return true, user, nil
}
