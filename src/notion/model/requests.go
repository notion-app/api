package model

type LoginRequest struct {
	AuthMethod  string `json:"auth_method"`
	AccessToken string `json:"access_token"`
}

// Lol @ the name
type SchoolRequestRequest struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

// This is the body for both creating and removing subscriptions
type SubscriptionRequest struct {
	NotebookId string `json:"notebook_id"`
	Name       string `json:"name"`
}

type AddSchoolRequest struct {
	SchoolId string `json:"school"`
}
