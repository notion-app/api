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

func NewUserResponse(dbu DbUser) UserResponse {
	return UserResponse{
		Id: dbu.Id,
		Name: dbu.Name,
		Email: dbu.Email,
		Verified: dbu.Verified,
		School: dbu.School.String,
		AuthMethod: dbu.AuthMethod,
		FbUserId: dbu.FbUserId,
		FbAuthToken: dbu.FbAuthToken,
		FbProfilePic: dbu.FbProfilePic,
	}
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

func NewNoteResponse(dbn DbNote) NoteResponse {
	nr := NoteResponse{
		Id: dbn.Id,
		Title: dbn.Title,
		Owner: dbn.Owner,
		CreatedAt: dbn.CreatedAt,
		UpdatedAt: dbn.UpdatedAt,
	}
	if len(dbn.Content) < NOTE_RESPONSE_LENGTH_LIMIT {
		nr.ContentPreview = dbn.Content
	} else {
		nr.ContentPreview = dbn.Content[:NOTE_RESPONSE_LENGTH_LIMIT]
	}
	return nr
}
