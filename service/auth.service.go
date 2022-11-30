package service

import (
	"context"
	"errors"
	"time"

	"emoney-service/models"
	"emoney-service/repository"
	"emoney-service/utils/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(ctx context.Context, username string, password string) (models.Auth, error)
	VerifyPassword(ctx context.Context, rawPassword string, hashedPassword string) error
}

type authService struct {
	DB         *gorm.DB
	Repository repository.UserRepository
}

func NewAuthService(db *gorm.DB, repository repository.UserRepository) AuthService {
	return &authService{
		DB:         db,
		Repository: repository,
	}
}

func (s *authService) Login(ctx context.Context, username string, password string) (models.Auth, error) {
	var result models.Auth
	db := s.DB
	user, err := s.Repository.FetchByUsername(ctx, db, username)
	if err != nil {
		return result, err
	}
	err = s.VerifyPassword(ctx, password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return result, errors.New("username or password mismatch")
	}
	token, err := token.GenerateToken(user.ID)
	if err != nil {
		return result, err
	}
	result.Token = token
	result.ExpiredAt = time.Now().Add(time.Hour * 1)
	return result, nil
}

func (s *authService) VerifyPassword(ctx context.Context, rawPassword string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
}
