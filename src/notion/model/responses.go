package model

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
