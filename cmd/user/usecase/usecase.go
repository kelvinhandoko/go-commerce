package usecase

import (
	"go-commerce/cmd/user/service"
	"go-commerce/models"
)

type UserUsecase struct {
	UserService service.UserService
}

func NewUserUsecase(userService service.UserService) *UserUsecase {
	return &UserUsecase{
		UserService: userService,
	}
}

func (uc *UserUsecase) GetUserByEmail(email string) (*models.User, error) {
	user, err := uc.UserService.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUsecase) CreateUser(user *models.User) error {
	_, err := uc.UserService.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
