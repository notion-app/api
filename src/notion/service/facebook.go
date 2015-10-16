package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"notion/log"
	"notion/model"
)

type Facebook struct{}

func (f Facebook) genericGet(authToken string, urlEndpoint string, st interface{}) (bool, interface{}, error) {
	resp, err := http.Get(fmt.Sprintf("https://graph.facebook.com/v2.4/%v?redirect=false&access_token=%v", urlEndpoint, authToken))
	if log.Error(err) {
		return false, st, err
	}
	if resp.StatusCode == 400 {
		return false, st, nil
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if log.Error(err) {
		return false, st, err
	}
	err = json.Unmarshal(bytes, &st)
	if log.Error(err) {
		return false, st, err
	}
	return true, st, nil
}

// Returns information about the currently logged in user
//  bool: whether or not the authToken is valid
//  struct: the current user information
//  error: these would always represent ISE's
func (f Facebook) GetCurrentUser(authToken string) (bool, model.FbCurrentUser, error) {
	var data model.FbCurrentUser
	in, st, err := f.genericGet(authToken, "me", &data)
	return in, *st.(*model.FbCurrentUser), err
}

// Calls a facebook api endpoint to make a short-lived token long-lived
// Returns a boolean as to whether or not the token provided is valid
func (f Facebook) ExtendToken(authToken string) (bool, error) {
	return false, nil
}

func (f Facebook) GetProfilePic(authToken string) (bool, model.FbProfilePic, error) {
	var data model.FbProfilePic
	in, st, err := f.genericGet(authToken, "me/picture", &data)
	return in, *st.(*model.FbProfilePic), err
}
