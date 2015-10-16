package model

type LoginResponse struct {
	UserId     string `json:"user_id"`
	Name       string `json:"name"`
	Token      string `json:"token"`
	ProfilePic string `json:"profile_pic"`
}

type AllSchoolsResponse struct {
	Schools []DbSchool `json:"schools"`
}

type UserSubscriptionsResponse struct {
	Subscriptions []DbSubscription `json:"subscriptions"`
}
