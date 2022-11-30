package service

import (
	"context"
	"emoney-service/models"
	"emoney-service/presentation"
	"emoney-service/repository"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"gorm.io/gorm"
)

type TransactionService interface {
	Confirm(ctx context.Context, userId uint, billerId uint) (bool, error)
	FetchInquiry(ctx context.Context) ([]models.Biller, error)
	FetchHistoryByUserID(ctx context.Context, id uint) ([]models.Transaction, error)
}

type transactionService struct {
	DB                       *gorm.DB
	TransactionRepository    repository.TransactionRepository
	BalanceRepository        repository.BalanceRepository
	BalanceHistoryRepository repository.BalanceHistoryRepository
	UserRepository           repository.UserRepository
}

func NewTransactionService(db *gorm.DB, transactionRepository repository.TransactionRepository, balanceRepository repository.BalanceRepository, balanceHistoryRepository repository.BalanceHistoryRepository, userRepository repository.UserRepository) TransactionService {
	return &transactionService{
		DB:                       db,
		TransactionRepository:    transactionRepository,
		BalanceRepository:        balanceRepository,
		BalanceHistoryRepository: balanceHistoryRepository,
		UserRepository:           userRepository,
	}
}

func (s *transactionService) Confirm(ctx context.Context, userId uint, billerId uint) (bool, error) {
	var bindingBiller presentation.BillferFetchDetailResponse
	urlBiller := fmt.Sprintf("https://phoenix-imkas.ottodigital.id/interview/biller/v1/detail?billerId=%v", billerId)
	fetchBiller, err := http.Get(urlBiller)
	if err != nil {
		return false, err
	}
	resBiller, err := ioutil.ReadAll(fetchBiller.Body)
	if err != nil {
		return false, err
	}
	json.Unmarshal(resBiller, &bindingBiller)
	if bindingBiller.Code != 200 {
		return false, errors.New("biller not found")
	}
	db := s.DB.Begin()
	user, err := s.UserRepository.FetchByUserID(ctx, db, userId)
	if err != nil {
		db.Rollback()
		return false, err
	}
	err = s.TransactionRepository.Create(ctx, db, bindingBiller.BillerData, user)
	if err != nil {
		db.Rollback()
		return false, err
	}
	resBalance, err := s.BalanceRepository.FetchByUserID(ctx, db, userId)
	if err != nil {
		db.Rollback()
		return false, err
	}
	newBalance := resBalance.Balance - bindingBiller.BillerData.Price - bindingBiller.BillerData.Fee
	if newBalance < 0 {
		db.Rollback()
		return false, errors.New("amount not enough")
	}
	err = s.BalanceRepository.Update(ctx, db, newBalance, userId)
	if err != nil {
		db.Rollback()
		return false, err
	}
	err = s.BalanceHistoryRepository.Create(ctx, db, newBalance*-1, userId)
	if err != nil {
		db.Rollback()
		return false, err
	}
	db.Commit()
	return true, nil
}

func (s *transactionService) FetchInquiry(ctx context.Context) ([]models.Biller, error) {
	var response []models.Biller
	var bindingBillers presentation.BillerFetchAllResponse
	fetchBillers, err := http.Get("https://phoenix-imkas.ottodigital.id/interview/biller/v1/list")
	if err != nil {
		return nil, err
	}
	resBillers, err := ioutil.ReadAll(fetchBillers.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resBillers, &bindingBillers)
	if bindingBillers.Code != 200 {
		return nil, errors.New("biller not found")
	}
	response = bindingBillers.BillerData
	return response, nil
}

func (s *transactionService) FetchHistoryByUserID(ctx context.Context, id uint) ([]models.Transaction, error) {
	res, err := s.TransactionRepository.FetchHistoryByUser(ctx, s.DB, id)
	if err != nil {
		return res, err
	}
	return res, nil
}
