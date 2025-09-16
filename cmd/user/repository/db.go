package repository

import (
	"context"
	"errors"
	"go-commerce/models"

	"gorm.io/gorm"
)

func (repo *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := repo.Database.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user, nil
		}

		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	var user models.User
	err := repo.Database.WithContext(ctx).Omit("password").Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user, nil
		}

		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) CreateUser(user *models.User) (int64, error) {
	err := repo.Database.Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}
