package model

type LoginRequest struct {
  AuthMethod string `json:"auth_method"`
  AccessToken string `json:"access_token"`
}

// Lol @ the name
type SchoolRequestRequest struct {
  Name string `json:"name"`
  Location string `json:"location"`
}
