package models

import "gorm.io/gorm"

type Biller struct {
	gorm.Model
	Category    string
	Product     string
	Description string
	Price       int
	Fee         int
}

type Transaction struct {
	gorm.Model
	UserID   uint
	BillerID uint
	User     User
	Biller   Biller
	Status   string
}
