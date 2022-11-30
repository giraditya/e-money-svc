package repository

import (
	"context"
	"emoney-service/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, db *gorm.DB, user models.User) (models.User, error)
	FetchByUsername(ctx context.Context, db *gorm.DB, username string) (models.User, error)
	FetchByUserID(ctx context.Context, db *gorm.DB, id uint) (models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(ctx context.Context, db *gorm.DB, user models.User) (models.User, error) {
	err := db.Create(&user).Error
	return user, err
}

func (r *userRepository) FetchByUsername(ctx context.Context, db *gorm.DB, username string) (models.User, error) {
	var user models.User
	err := db.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *userRepository) FetchByUserID(ctx context.Context, db *gorm.DB, id uint) (models.User, error) {
	var user models.User
	err := db.Where("id = ?", id).First(&user).Error
	return user, err
}
