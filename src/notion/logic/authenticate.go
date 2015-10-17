package logic

import (
	"notion/db"
	"notion/errors"
	"notion/model"
)

// Takes in an auth token and returns the user who is associated with that auth
// token or a user-servicible error if the auth token is invalid or
// the database connection fails
func AuthenticateNotionUser(token string) (model.DbUser, error) {
	in, user, err := db.GetUserByToken(token)
	if !in {
		return user, errors.Unauthorized("notion")
	}
	return user, err
}
