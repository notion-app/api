package model

type LoginRequest struct {
  AuthMethod string `json:"auth_method"`
  AccessToken string `json:"access_token"`
}
