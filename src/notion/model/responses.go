package model

type UserResponse struct {
	Id           string         `json:"id"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Verified     bool           `json:"verified"`
	School       string 				`json:"school"`
	AuthMethod   string         `json:"auth_method"`
	FbUserId     string         `json:"fb_user_id"`
	FbAuthToken  string         `json:"fb_auth_token"`
	FbProfilePic string         `json:"fb_profile_pic"`
}

type LoginResponse struct {
	UserId     string `json:"user_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	ProfilePic string `json:"profile_pic"`
}

type AllSchoolsResponse struct {
	Schools []DbSchool `json:"schools"`
}

type UserSubscriptionsResponse struct {
	Subscriptions []DbSubscription `json:"subscriptions"`
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
