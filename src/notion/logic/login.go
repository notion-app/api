package logic

import (
	"net/http"
	"notion/db"
	"notion/log"
	"notion/model"
	"notion/service"
	"notion/util"
)

func DoUserCreateOrLogin(lrq model.LoginRequest) (int, model.UserResponse, error) {
	var userResponse model.UserResponse
	var returnCode int
	fbUser, err := service.Facebook{}.GetCurrentUser(lrq.AccessToken)
	if log.Error(err) {
		return returnCode, userResponse, err
	}
	fbPicture, err := service.Facebook{}.GetProfilePic(lrq.AccessToken)
	if log.Error(err) {
		return returnCode, userResponse, err
	}
	in, dbUser, err := db.GetUserByFacebookId(fbUser.Id)
	if in {
		err = DoUserLogin(lrq, &dbUser, fbPicture)
		returnCode = http.StatusAccepted
	} else {
		dbUser, err = DoFbUserCreate(lrq, fbUser, fbPicture)
		returnCode = http.StatusCreated
	}
	if log.Error(err) {
		return returnCode, userResponse, err
	}
	userResponse.FromDb(dbUser)
	return returnCode, userResponse, nil
}

func DoFbUserCreate(lrq model.LoginRequest, fbUser model.FbCurrentUser, fbPicture model.FbProfilePic) (model.DbUser, error) {
	user := model.DbUser{
		Id:           util.NewId(),
		Name:         fbUser.Name,
		Email:        fbUser.Email,
		Verified:     false,
		AuthMethod:   lrq.AuthMethod,
		FbUserId:     fbUser.Id,
		FbAuthToken:  lrq.AccessToken,
		FbProfilePic: fbPicture.Data.Url,
	}
	return user, db.CreateUser(user)
}

func DoUserLogin(lrq model.LoginRequest, u *model.DbUser, fbPicture model.FbProfilePic) error {
	u.FbAuthToken = lrq.AccessToken
	u.FbProfilePic = fbPicture.Data.Url
	return db.UpdateUser(*u)
}
