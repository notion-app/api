package db

import (
	"notion/log"
	"notion/model"
)

// Gets a user by their notion-assigned Id
// Returns whether the user exists, the user model, and an error
func GetUserById(id string) (bool, model.DbUser, error) {
	user := model.DbUser{}
	in, err := GenericGetOne(&user, "select * from users where id=$1", id)
	return in, user, err
}

func GetUserByFacebookId(facebookId string) (bool, model.DbUser, error) {
	user := model.DbUser{}
	in, err := GenericGetOne(&user, "select * from users where fb_user_id=$1", facebookId)
	return in, user, err
}

func GetUserByToken(token string) (bool, model.DbUser, error) {
	user := model.DbUser{}
	in, err := GenericGetOne(&user, "select * from users where fb_token=$1", token)
	return in, user, err
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
