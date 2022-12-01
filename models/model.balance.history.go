package models

import "gorm.io/gorm"

type BalanceHistory struct {
	gorm.Model
	User   User
	UserID uint
	Amount int
}
