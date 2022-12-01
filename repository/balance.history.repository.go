package repository

import (
	"context"
	"emoney-service/models"

	"gorm.io/gorm"
)

type BalanceHistoryRepository interface {
	Create(ctx context.Context, db *gorm.DB, amount int, userid uint) error
}

type balanceHistoryRepository struct{}

func NewBalanceHistoryRepository() BalanceHistoryRepository {
	return &balanceHistoryRepository{}
}

func (r *balanceHistoryRepository) Create(ctx context.Context, db *gorm.DB, amount int, userid uint) error {
	var balanceHistory = models.BalanceHistory{
		UserID: userid,
		Amount: amount,
	}
	err := db.Create(&balanceHistory).Error
	if err != nil {
		return err
	}
	return nil
}
