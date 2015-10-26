package model

type LoginRequest struct {
	AuthMethod  string `json:"auth_method" binding:"required"`
	AccessToken string `json:"access_token" binding:"required"`
}

// Lol @ the name
type SchoolRequestRequest struct {
	Name     string `json:"name" binding:"required"`
	Location string `json:"location" binding:"required"`
}

// This is the body for both creating and removing subscriptions
type SubscriptionRequest struct {
	NotebookId string `json:"notebook_id" binding:"required"`
	Name       string `json:"name"`
}

type ModifySchoolRequest struct {
	SchoolId string `json:"school" binding:"required"`
}
