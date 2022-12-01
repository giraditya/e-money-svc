package models

import (
	"time"

	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model
	Email     string
	Username  string
	Password  string
	Token     string
	ExpiredAt time.Time
}
