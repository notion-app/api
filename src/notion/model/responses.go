package model

type UserResponse struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Verified     bool   `json:"verified"`
	School       string `json:"school"`
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
