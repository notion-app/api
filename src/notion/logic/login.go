
package logic

import (
  "notion/errors"
  "notion/db"
  "notion/log"
  "notion/model"
  "notion/service"
  "notion/util"
)

func DoUserCreateOrLogin(lrq model.LoginRequest) error {
  valid, fbUser, err := service.Facebook{}.GetCurrentUser(lrq.AccessToken)
  if log.Error(err) {
    return err
  }
  if !valid {
    return errors.Unauthorized("facebook")
  }
  valid, fbPicture, err := service.Facebook{}.GetProfilePic(lrq.AccessToken)
  if log.Error(err) {
    return err
  }
  if !valid {
    return errors.Unauthorized("facebook")
  }
  in, dbUser, err := db.GetUserByFacebookId(fbUser.Id)
  if in {
    return DoUserLogin(lrq, dbUser, fbPicture)
  } else {
    return DoFbUserCreate(lrq, fbUser, fbPicture)
  }
}

func DoFbUserCreate(lrq model.LoginRequest, fbUser model.FbCurrentUser, fbPicture model.FbProfilePic) error {
  user := model.DbUser{
    Id: util.NewId(),
    Name: fbUser.Name,
    Verified: false,
    AuthMethod: lrq.AuthMethod,
    FbUserId: fbUser.Id,
    FbAuthToken: lrq.AccessToken,
    FbProfilePic: fbPicture.Url,
  }
  return db.CreateUser(user)
}

func DoUserLogin(lrq model.LoginRequest, u model.DbUser, fbPicture model.FbProfilePic) error {
  u.FbAuthToken = lrq.AccessToken
  u.FbProfilePic = fbPicture.Url
  return db.UpdateUser(u)
}
