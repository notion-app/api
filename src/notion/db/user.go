
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
  err := dbmap.SelectOne(&user, "select * from users where id=$1", id)
  if err != nil {
    switch err {
    case sql.ErrNoRows:
      return false, user, nil
    default:
      log.Error(err)
      return false, user, err
    }
  }
  return true, user, nil
}

func GetUserByFacebookId(facebookId string) (bool, model.DbUser, error) {
  log.Info("Getting user by facebook id " + facebookId)
  var user model.DbUser
  err := dbmap.SelectOne(&user, "select * from users where fb_user_id=$1", facebookId)
  if err != nil {
    switch err {
    case sql.ErrNoRows:
      return false, user, nil
    default:
      log.Error(err)
      return false, user, err
    }
  }
  return true, user, nil
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
