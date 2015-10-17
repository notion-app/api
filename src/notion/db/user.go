package db

import (
	"notion/errors"
	"notion/log"
	"notion/model"
)

// Gets a user by their notion-assigned Id
func GetUserById(id string) (model.DbUser, error) {
	var user model.DbUser
	err := GenericGetOne("users", "id", id, &user)
	return user, err
}

func GetUserByFacebookId(facebookId string) (model.DbUser, error) {
	var user model.DbUser
	err := GenericGetOne("users", "fb_user_id", facebookId, &user)
	return user, err
}

func GetUserByToken(token string) (model.DbUser, error) {
	var user model.DbUser
	err := GenericGetOne("users", "fb_auth_token", token, &user)
	return user, err
}

func CreateUser(u model.DbUser) error {
	log.Info("Creating new user " + u.Id)
	err := dbmap.Insert(&u)
	if log.Error(err) {
		return errors.ISE()
	}
	return nil
}

func UpdateUser(user model.DbUser) error {
	log.Info("Updating user token and profile pic")
	_, err := dbmap.Update(&user)
	if log.Error(err) {
		return errors.ISE()
	}
	return nil
}
