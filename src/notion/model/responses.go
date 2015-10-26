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

// This function serializes the sql.NullString field of School so it returns
// a raw string, and the clients are expected to know that it might be an
// empty string
func NewUserResponse(dbu DbUser) UserResponse {
	return UserResponse{
		Id:           dbu.Id,
		Name:         dbu.Name,
		Email:        dbu.Email,
		Verified:     dbu.Verified,
		School:       dbu.School.String,
		AuthMethod:   dbu.AuthMethod,
		FbUserId:     dbu.FbUserId,
		FbAuthToken:  dbu.FbAuthToken,
		FbProfilePic: dbu.FbProfilePic,
	}
}

type CourseResponse struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Number string `json:"number"`
}

func CourseResponseWithoutSchool(dbc DbCourse) CourseResponse {
	return CourseResponse{
		Id: dbc.Id,
		Name: dbc.Name,
		Number: dbc.Number,
	}
}

type SectionResponse struct {
	Id string `json:"id"`
	NotebookId string `json:"notebook_id"`
	Crn string `json:"crn"`
	Professor string `json:"professor"`
	Year string `json:"year"`
	Semester string `json:"semester"`
	Time string `json:"time"`
	Verified bool `json:"verified"`
}

func SectionResponseWithoutCourse(dbc DbCourseSection) SectionResponse {
	return SectionResponse{
		Id: dbc.Id,
		NotebookId: dbc.NotebookId,
		Crn: dbc.Crn,
		Professor: dbc.Professor,
		Year: dbc.Year,
		Semester: dbc.Semester,
		Time: dbc.Time,
		Verified: dbc.Verified,
	}
}

type TopicResponse struct {
	Id    string         `json:"id"`
	Notes []NoteResponse `json:"notes"`
}

type NoteResponse struct {
	Id             string     `json:"id"`
	Title          string     `json:"title"`
	Owner          string     `json:"owner"`
	ContentPreview string     `json:"content_preview"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

func NewNoteResponse(dbn DbNote) NoteResponse {
	nr := NoteResponse{
		Id:        dbn.Id,
		Title:     dbn.Title.String,
		Owner:     dbn.Owner,
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

type SubscriptionResponse struct {
	UserId string `json:"id"`
	NotebookId string `json:"notebook_id"`
	Name string `json:"name"`
}

func NewSubscriptionResponse(dbs DbSubscription) SubscriptionResponse {
	return SubscriptionResponse{
		UserId: dbs.UserId,
		NotebookId: dbs.NotebookId,
		Name: dbs.Name.String,
	}
}
