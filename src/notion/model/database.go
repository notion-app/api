package model

type DbSchool struct {
	Id       string `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Location string `db:"location" json:"location"`
}

type DbSchoolRequest struct {
	Id              string `db:"id" json:"id"`
	RequesterUserId string `db:"requester_user_id" json:"requester_user_id"`
	Name            string `db:"name" json:"name"`
	Location        string `db:"location" json:"location"`
}

type DbCourse struct {
	Id       string `db:"id" json:"id"`
	SchoolId string `db:"school_id" json:"school_id"`
	Name     string `db:"name" json:"name"`
	Number   string `db:"number" json:"number"`
}

type DbCourseSection struct {
	Id        string `db:"id" json:"id"`
	CourseId  string `db:"course_id" json:"course_id"`
	Crn       string `db:"crn" json:"crn"`
	Professor string `db:"professor" json:"professor"`
	Year      string `db:"year" json:"year"`
	Semester  string `db:"semester" json:"semester"`
	Time      string `db:"time" json:"time"`
	Verified  bool   `db:"verified" json:"verified"`
}

type DbUser struct {
	Id           string `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	Verified     bool   `db:"verified" json:"verified"`
	AuthMethod   string `db:"auth_method" json:"auth_method"`
	FbUserId     string `db:"fb_user_id" json:"fb_user_id"`
	FbAuthToken  string `db:"fb_auth_token" json:"fb_auth_token"`
	FbProfilePic string `db:"fb_profile_pic" json:"fb_profile_pic"`
}

type DbNotebook struct {
	Id        string `db:"id" json:"id"`
	SectionId string `db:"section_id" json:"section_id"`
	Name      string `db:"name" json:"name"`
	Owner     string `db:"owner" json:"owner"`
	Privacy   string `db:"privacy" json:"privacy"`
}

type DbTopic struct {
	Id         string `db:"id" json:"id"`
	NotebookId string `db:"notebook_id" json:"notebook_id"`
}

type DbNote struct {
	Id      string `db:"id" json:"id"`
	TopicId string `db:"topic_id" json:"topic_id"`
	Name    string `db:"name" json:"name"`
	Owner   string `db:"owner" json:"owner"`
	Content string `db:"content" json:"content"`
}

type DbSubscription struct {
	UserId     string `db:"user_id" json:"user_id"`
	NotebookId string `db:"notebook_id" json:"notebook_id"`
}
