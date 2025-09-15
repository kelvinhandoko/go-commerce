package service

import (
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

func (svc *UserService) GetUserByEmail(email string) (*models.User, error) {
	user, err := svc.UserRepo.GetUserByEmail(email)
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
