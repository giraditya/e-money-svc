package models

import "gorm.io/gorm"

type Biller struct {
	gorm.Model
	Category    string `json:"category"`
	Product     string `json:"product"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Fee         int    `json:"fee"`
}

type Transaction struct {
	gorm.Model
	UserID   uint `json:"user_id"`
	BillerID uint `json:"biller"`
	User     User
	Biller   Biller
	Status   string `json:"status"`
}
