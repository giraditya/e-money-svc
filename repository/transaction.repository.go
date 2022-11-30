package repository

import (
	"context"
	"emoney-service/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(ctx context.Context, db *gorm.DB, biller models.Biller, user models.User) error
	FetchHistoryByUser(ctx context.Context, db *gorm.DB, userID uint) ([]models.Transaction, error)
}

type transactionRepository struct{}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{}
}

func (r *transactionRepository) Create(ctx context.Context, db *gorm.DB, biller models.Biller, user models.User) error {
	err := db.Create(&biller).Error
	if err != nil {
		return err
	}
	err = db.Create(&models.Transaction{
		UserID:   user.ID,
		BillerID: biller.ID,
		Status:   "OK",
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *transactionRepository) FetchHistoryByUser(ctx context.Context, db *gorm.DB, userID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := db.Model(&models.Transaction{}).Preload("User").Preload("Biller").Where("user_id = ?", userID).Find(&transactions).Error
	return transactions, err
}
