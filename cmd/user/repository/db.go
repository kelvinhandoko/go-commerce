package repository

import "go-commerce/models"

func (repo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := repo.Database.Where("email = ?", email).First(&user).Error
	if err != nil {
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
