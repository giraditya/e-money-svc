package service

import (
	"context"
	"emoney-service/models"
	"emoney-service/repository"

	"gorm.io/gorm"
)

type BalanceService interface {
	FetchByUserID(ctx context.Context, id uint) (models.Balance, error)
}

type balanceService struct {
	DB         *gorm.DB
	Repository repository.BalanceRepository
}

func NewBalanceService(db *gorm.DB, repository repository.BalanceRepository) BalanceService {
	return &balanceService{
		DB:         db,
		Repository: repository,
	}
}

func (s *balanceService) FetchByUserID(ctx context.Context, id uint) (models.Balance, error) {
	db := s.DB
	res, err := s.Repository.FetchByUserID(ctx, db, id)
	if err != nil {
		return models.Balance{}, err
	} else {
		return models.Balance{
			User:    res.User,
			UserID:  id,
			Balance: res.Balance,
		}, nil
	}
}
