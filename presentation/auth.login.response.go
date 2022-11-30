package presentation

import "time"

type AuthLoginResponse struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expiredAt"`
}
