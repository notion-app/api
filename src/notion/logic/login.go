
package logic

import (
  "net/http"
  "notion/errors"
  "notion/db"
  "notion/log"
  "notion/model"
  "notion/service"
  "notion/util"
)

func DoUserCreateOrLogin(lrq model.LoginRequest) (int, model.LoginResponse, error) {
  var loginResponse model.LoginResponse
  var returnCode int
  valid, fbUser, err := service.Facebook{}.GetCurrentUser(lrq.AccessToken)
  if log.Error(err) {
    return returnCode, loginResponse, err
  }
  if !valid {
    return returnCode, loginResponse, errors.Unauthorized("facebook")
  }
  valid, fbPicture, err := service.Facebook{}.GetProfilePic(lrq.AccessToken)
  if log.Error(err) {
    return returnCode, loginResponse, err
  }
  if !valid {
    return returnCode, loginResponse, errors.Unauthorized("facebook")
  }
  in, dbUser, err := db.GetUserByFacebookId(fbUser.Id)
  if in {
    err = DoUserLogin(lrq, dbUser, fbPicture)
    returnCode = http.StatusAccepted
  } else {
    dbUser, err = DoFbUserCreate(lrq, fbUser, fbPicture)
    returnCode = http.StatusCreated
  }
  if log.Error(err) {
    return returnCode, loginResponse, err
  }
  loginResponse.UserId = dbUser.Id
  loginResponse.Name = dbUser.Name
  loginResponse.Token = dbUser.FbAuthToken
  loginResponse.ProfilePic = dbUser.FbProfilePic
  return returnCode, loginResponse, nil
}

func DoFbUserCreate(lrq model.LoginRequest, fbUser model.FbCurrentUser, fbPicture model.FbProfilePic) (model.DbUser, error) {
  user := model.DbUser{
    Id: util.NewId(),
    Name: fbUser.Name,
    Verified: false,
    AuthMethod: lrq.AuthMethod,
    FbUserId: fbUser.Id,
    FbAuthToken: lrq.AccessToken,
    FbProfilePic: fbPicture.Data.Url,
  }
  return user, db.CreateUser(user)
}

func DoUserLogin(lrq model.LoginRequest, u model.DbUser, fbPicture model.FbProfilePic) error {
  u.FbAuthToken = lrq.AccessToken
  u.FbProfilePic = fbPicture.Data.Url
  return db.UpdateUser(u)
}
