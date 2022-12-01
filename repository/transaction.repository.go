package repository

import (
	"context"
	"emoney-service/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(ctx context.Context, db *gorm.DB, biller models.Biller, user models.User) error
	FetchHistoryByUser(ctx context.Context, db *gorm.DB, userid uint) ([]models.Transaction, error)
}

type transactionRepository struct{}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{}
}

func (r *transactionRepository) Create(ctx context.Context, db *gorm.DB, biller models.Biller, user models.User) error {
	var transactions = models.Transaction{
		UserID:   user.ID,
		BillerID: biller.ID,
		Status:   "OK",
	}
	err := db.Create(&biller)
	if err.Error != nil {
		return err.Error
	}
	err = db.Create(&transactions)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *transactionRepository) FetchHistoryByUser(ctx context.Context, db *gorm.DB, userid uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := db.Model(&transactions).Preload("User").Preload("Biller").Where("user_id = ?", userid).Find(&transactions).Error
	return transactions, err
}
