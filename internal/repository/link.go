package repository

import (
	"bitly-clone/models"

	"gorm.io/gorm"
)

type LinkRepository struct {
	DB *gorm.DB
}

func (repo *Repositories) Link() *LinkRepository {
	return &LinkRepository{DB: repo.DB}
}

func (lr *LinkRepository) FindByShortURL(userID *uint, url string, link *models.Link) {
	q := lr.DB.
		Model(&models.Link{})

	if userID != nil {
		q = q.Where("user_id = ?", *userID)
	}

	q.Where("short_url = ?", url).
		Find(link)
}

func (lr *LinkRepository) FindByURL(userID *uint, url string, link *models.Link) {
	q := lr.DB.
		Model(&models.Link{})

	if userID != nil {
		q = q.Where("user_id = ?", *userID)
	}

	q.Where("url = ?", url).
		Find(link)
}

func (lr *LinkRepository) FindByID(userID *uint, ID uint, link *models.Link) {

	q := lr.DB.
		Model(&models.Link{})

	if userID != nil {
		q = q.Where("user_id = ?", *userID)
	}

	q.Where("id = ?", ID).
		Find(link)
}

func (lr *LinkRepository) Create(link *models.Link) error {
	return lr.DB.Create(link).Error
}

func (lr *LinkRepository) Delete(link *models.Link) error {
	return lr.DB.Delete(link).Error
}

func (lr *LinkRepository) List(userID uint) *gorm.DB {
	return lr.DB.
		Model(&models.Link{}).
		Where("user_id = ?", userID)
}
