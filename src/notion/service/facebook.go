package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"notion/log"
	"notion/model"
)

type Facebook struct{}

const (
	FacebookBaseUrl = "https://graph.facebook.com/v2.4/"
)

func (f Facebook) genericGet(endpoint string, struc interface{}, token string, params url.Values) error {
	params.Set("access_token", token)
	fullURL := fmt.Sprintf("%v%v?%v", FacebookBaseUrl, endpoint, params.Encode())
	resp, err := http.Get(fullURL)
	if log.Error(err) {
		return fmt.Errorf("ISE")
	}
	if resp.StatusCode == 400 {
		return fmt.Errorf("Facebook unauthorized")
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if log.Error(err) {
		return fmt.Errorf("ISE")
	}
	err = json.Unmarshal(bytes, &struc)
	if log.Error(err) {
		return fmt.Errorf("ISE")
	}
	return nil
}

// Returns information about the currently logged in user
//  bool: whether or not the authToken is valid
//  struct: the current user information
//  error: these would always represent ISE's
func (f Facebook) GetCurrentUser(authToken string) (model.FbCurrentUser, error) {
	var data model.FbCurrentUser
	params := url.Values{}
	params.Set("fields", "id,name,email")
	err := f.genericGet("me", &data, authToken, params)
	return data, err
}

func (f Facebook) GetProfilePic(authToken string) (model.FbProfilePic, error) {
	var data model.FbProfilePic
	params := url.Values{}
	params.Set("type", "large")
	params.Set("redirect", "false")
	err := f.genericGet("me/picture", &data, authToken, params)
	return data, err
}

// Calls a facebook api endpoint to make a short-lived token long-lived
// Returns a boolean as to whether or not the token provided is valid
func (f Facebook) ExtendToken(authToken string) (bool, error) {
	return false, nil
}
