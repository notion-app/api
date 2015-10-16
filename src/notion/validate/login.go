package validate

import (
	"notion/errors"
	"notion/model"
)

func Login(request model.LoginRequest) error {
	if request.AuthMethod == "facebook" {
		if request.AccessToken == "" {
			return errors.BadRequest("Must provide a facebook access token")
		}
	} else {
		return errors.BadRequest("Must provide an auth_method of 'facebook'")
	}
	return nil
}
