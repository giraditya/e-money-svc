package models

import "gorm.io/gorm"

type Balance struct {
	gorm.Model
	User    User
	UserID  uint
	Balance int `json:"balance"`
}
