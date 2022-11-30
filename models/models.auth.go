package models

import (
	"time"

	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expiredAt"`
}
