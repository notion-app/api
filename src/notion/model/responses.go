package model

import (
	"time"
)

const (
	NOTE_RESPONSE_LENGTH_LIMIT = 35
)

type UserResponse struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Verified     bool   `json:"verified"`
	School       string `json:"school_id"`
	AuthMethod   string `json:"auth_method"`
	FbUserId     string `json:"fb_user_id"`
	FbAuthToken  string `json:"fb_auth_token"`
	FbProfilePic string `json:"fb_profile_pic"`
}

func (u *UserResponse) FromDb(dbu DbUser) {
	u.Id = dbu.Id
	u.Name = dbu.Name
	u.Email = dbu.Email
	u.Verified = dbu.Verified
	u.School = dbu.School.String
	u.AuthMethod = dbu.AuthMethod
	u.FbUserId = dbu.FbUserId
	u.FbAuthToken = dbu.FbAuthToken
	u.FbProfilePic = dbu.FbProfilePic
}

type CoursesForSchoolResponse struct {
	Courses []CourseResponse `json:"courses"`
}

type CourseResponse struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Number string `json:"number"`
}

type SectionsForCourseResponse struct {
	Sections []DbCourseSection `json:"sections"`
}

type TopicResponse struct {
	Id string `json:"id"`
	Notes []NoteResponse `json:"notes"`
}

type NoteResponse struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Owner string `json:"owner"`
	ContentPreview string `json:"content_preview"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (nr *NoteResponse) FromDb(dbn DbNote) {
	nr.Id = dbn.Id
	nr.Title = dbn.Title
	nr.Owner = dbn.Owner
	nr.CreatedAt = dbn.CreatedAt
	nr.UpdatedAt = dbn.UpdatedAt
	if len(dbn.Content) < NOTE_RESPONSE_LENGTH_LIMIT {
		nr.ContentPreview = dbn.Content
	} else {
		nr.ContentPreview = dbn.Content[:NOTE_RESPONSE_LENGTH_LIMIT]
	}
}
