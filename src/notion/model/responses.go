package model

type LoginResponse struct {
  UserId string `json:"user_id"`
  Token string `json:"token"`
}
