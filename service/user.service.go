package service

import (
	"context"
	"emoney-service/models"
	"emoney-service/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Register(ctx context.Context, email string, username string, password string) (models.User, error)
	FetchByUsername(ctx context.Context, username string) (models.User, error)
	FetchByUserID(ctx context.Context, id uint) (models.User, error)
}

type userService struct {
	DB         *gorm.DB
	Repository repository.UserRepository
}

func NewUserService(db *gorm.DB, repository repository.UserRepository) UserService {
	return &userService{
		DB:         db,
		Repository: repository,
	}
}

func (s *userService) Register(ctx context.Context, email string, username string, password string) (models.User, error) {
	var db *gorm.DB = s.DB.Begin()

	userExist, _ := s.Repository.FetchByUsername(ctx, db, username)
	if userExist != (models.User{}) {
		return models.User{}, errors.New("user already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Email:    email,
		Username: username,
		Password: string(hashedPassword),
	}

	_, err = s.Repository.Create(ctx, db, user)
	if err != nil {
		db.Rollback()
		return models.User{}, err
	} else {
		db.Commit()
		return models.User{
			Email:    email,
			Username: username,
			Password: password,
		}, nil
	}
}

func (s *userService) FetchByUsername(ctx context.Context, username string) (models.User, error) {
	var db *gorm.DB = s.DB

	res, err := s.Repository.FetchByUsername(ctx, db, username)
	if err != nil {
		return models.User{}, err
	} else {
		return models.User{
			Email:    res.Email,
			Username: res.Username,
		}, nil
	}
}

func (s *userService) FetchByUserID(ctx context.Context, id uint) (models.User, error) {
	var db *gorm.DB = s.DB

	res, err := s.Repository.FetchByUserID(ctx, db, id)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		Model: gorm.Model{
			ID: res.ID,
		},
		Email:    res.Email,
		Username: res.Username,
		Password: "",
	}, nil
}
