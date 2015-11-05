package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notion/db"
	"notion/errors"
	"notion/log"
	"notion/model"
	"notion/service"
	"notion/util"
)

func Login(c *gin.Context) {
	var returnCode int
	var request model.LoginRequest

	// Parse the user request
	err := c.BindJSON(&request)
	if log.Error(err) {
		c.Error(err)
		return
	}

	// Right now we assume that the user is logging in with Facebook
	fbUser, err := service.Facebook{}.GetCurrentUser(request.AccessToken)
	if log.Error(err) {
		c.Error(errors.NewHttp(errors.ISE, "Error contacting facebook api"))
		return
	}

	// Then get the user's profile picture
	fbPicture, err := service.Facebook{}.GetProfilePic(request.AccessToken)
	if log.Error(err) {
		c.Error(errors.NewHttp(errors.ISE, "Error contacting facebook api"))
		return
	}

	// See if the user is already a notion user
	in, dbUser, err := db.GetUserByFacebookId(fbUser.Id)
	if log.Error(err) {
		c.Error(errors.NewISE())
		return
	}

	// If they are in the database, we just update their auth token
	if in {
		returnCode = http.StatusAccepted
		dbUser.FbAuthToken = request.AccessToken
		dbUser.FbProfilePic = fbPicture.Data.Url
		err = db.UpdateUser(dbUser)
	} else {
		returnCode = http.StatusCreated
		dbUser = model.DbUser{
			Id:           util.NewId(),
			Name:         fbUser.Name,
			Email:        fbUser.Email,
			Verified:     false,
			AuthMethod:   request.AuthMethod,
			FbUserId:     fbUser.Id,
			FbAuthToken:  request.AccessToken,
			FbProfilePic: fbPicture.Data.Url,
		}
		err = db.CreateUser(dbUser)
	}

	// Error check 'er yo
	if log.Error(err) {
		c.Error(errors.NewISE())
		return
	}

	// Throw back the user object at the requester
	c.JSON(returnCode, model.NewUserResponse(dbUser))

}
