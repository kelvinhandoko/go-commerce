package service

import (
	"context"
	"go-commerce/cmd/user/repository"
	"go-commerce/models"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (svc *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := svc.UserRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (svc *UserService) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	user, err := svc.UserRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (svc *UserService) CreateUser(user *models.User) (int64, error) {
	id, err := svc.UserRepo.CreateUser(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}
