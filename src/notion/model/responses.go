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
	Username     string `json:"username"`
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
	u := UserResponse{
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
	if dbu.Username.Valid {
		u.Username = dbu.Username.String
	} else {
		u.Username = ""
	}
	return u
}

type CourseResponse struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Number string `json:"number"`
}

func CourseResponseWithoutSchool(dbc DbCourse) CourseResponse {
	return CourseResponse{
		Id:     dbc.Id,
		Name:   dbc.Name,
		Number: dbc.Number,
	}
}

type SectionResponse struct {
	Id         string `json:"id"`
	NotebookId string `json:"notebook_id"`
	Crn        string `json:"crn"`
	Professor  string `json:"professor"`
	Year       string `json:"year"`
	Semester   string `json:"semester"`
	Time       string `json:"time"`
	Verified   bool   `json:"verified"`
}

func SectionResponseWithoutCourse(dbc DbCourseSection) SectionResponse {
	return SectionResponse{
		Id:         dbc.Id,
		NotebookId: dbc.NotebookId,
		Crn:        dbc.Crn,
		Professor:  dbc.Professor,
		Year:       dbc.Year,
		Semester:   dbc.Semester,
		Time:       dbc.Time,
		Verified:   dbc.Verified,
	}
}

type TopicResponse struct {
	Id    string             `json:"id"`
	Notes []FullNoteResponse `json:"notes"`
}

type FullNoteResponse struct {
	Id        string     `json:"id"`
	TopicId   string     `json:"topic_id"`
	Title     string     `json:"title"`
	Owner     string     `json:"owner"`
	Content   string     `json:"content"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func NewFullNoteResponse(dbn DbNote) FullNoteResponse {
	nr := FullNoteResponse{
		Id:        dbn.Id,
		TopicId:   dbn.TopicId,
		Title:     dbn.Title.String,
		Owner:     dbn.Owner,
		Content:   dbn.Content,
		CreatedAt: dbn.CreatedAt,
		UpdatedAt: dbn.UpdatedAt,
	}
	return nr
}

type SubscriptionResponse struct {
	Id         string          `json:"id"`
	UserId     string          `json:"user_id"`
	NotebookId string          `json:"notebook_id"`
	Name       string          `json:"name"`
	Course     DbCourse        `json:"course"`
	Section    DbCourseSection `json:"section"`
}

func NewSubscriptionResponse(dbs DbSubscription, course DbCourse, section DbCourseSection) SubscriptionResponse {
	return SubscriptionResponse{
		Id:         dbs.UserId + dbs.NotebookId,
		UserId:     dbs.UserId,
		NotebookId: dbs.NotebookId,
		Name:       dbs.Name.String,
		Course:     course,
		Section:    section,
	}
}
