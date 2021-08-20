package repository

import (
	"bitly-clone/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (Repo *Repositories) User() *UserRepository {
	return &UserRepository{DB: Repo.DB}
}

func (ur *UserRepository) FindByUsername(username string, user *models.User) {
	ur.DB.
		Model(&models.User{}).
		Where("username = ?", username).
		Scan(user)
}

func (ur *UserRepository) FindByToken(token string, user *models.User) {
	ur.DB.Model(&models.User{}).Where("token = ?", token).First(&user)
}

func (ur *UserRepository) Create(user *models.User) error {
	return ur.DB.Create(user).Error
}
