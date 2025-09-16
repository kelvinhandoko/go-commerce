package usecase

import (
	"context"
	"errors"
	"go-commerce/cmd/user/service"
	"go-commerce/infrastucture/log"
	"go-commerce/models"
	"go-commerce/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type UserUsecase struct {
	UserService service.UserService
	JWTSecret   string
}

func NewUserUsecase(userService service.UserService, jwtSecret string) *UserUsecase {
	return &UserUsecase{
		UserService: userService,
		JWTSecret:   jwtSecret,
	}
}

func (uc *UserUsecase) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := uc.UserService.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUsecase) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	user, err := uc.UserService.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUsecase) CreateUser(user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"email": user.Email,
		}).Errorf("failed to hash password: %v", err)
		return err
	}
	user.Password = hashedPassword

	_, err = uc.UserService.CreateUser(user)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"email": user.Email,
			"name":  user.Name,
		}).Errorf("failed to create user: %v", err)
		return err
	}
	return nil
}

// Login login by given param pointer of models.LoginParameter.
//
// It returns string, and nil error when successful.
// Otherwise, empty string, and error will be returned.
func (uc *UserUsecase) Login(ctx context.Context, param models.LoginParameter) (string, error) {
	user, err := uc.UserService.GetUserByEmail(ctx, param.Email)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"email": param.Email,
		}).Errorf("failed to get user by email: %v", err)
		return "", err
	}

	if user.ID == 0 {
		return "", errors.New("email not registered")
	}

	checkPass, err := utils.VerifyPassword(user.Password, param.Password)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"email": param.Email,
		}).Errorf("failed to verify password: %v", err)
		return "", err
	}

	if !checkPass {
		return "", errors.New("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, err := token.SignedString([]byte(uc.JWTSecret))

	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"email": param.Email,
		}).Errorf("failed to sign token: %v", err)
		return "", err
	}

	return tokenString, nil

}
