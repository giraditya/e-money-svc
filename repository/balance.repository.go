package repository

import (
	"context"
	"emoney-service/models"
	"errors"

	"gorm.io/gorm"
)

type BalanceRepository interface {
	Decrease(ctx context.Context, db *gorm.DB, amount int, userid uint) error
	FetchByUserID(ctx context.Context, db *gorm.DB, userID uint) (models.Balance, error)
}

type balanceRepository struct{}

func NewBalanceRepository() BalanceRepository {
	return &balanceRepository{}
}

func (r *balanceRepository) Decrease(ctx context.Context, db *gorm.DB, amount int, userid uint) error {
	var balance models.Balance
	var newAmount int
	err := db.Where("user_id = ?", userid).First(&balance).Error
	if err != nil {
		return err
	}

	if newAmount = balance.Balance - amount; newAmount < 0 {
		return errors.New("balance not enough to decrease")
	}

	err = db.Model(&models.Balance{}).Where("user_id = ?", userid).Update("balance", newAmount).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *balanceRepository) FetchByUserID(ctx context.Context, db *gorm.DB, userID uint) (models.Balance, error) {
	var balance models.Balance
	err := db.Where("user_id =?", userID).Preload("User").Find(&balance).Error
	return balance, err
}
