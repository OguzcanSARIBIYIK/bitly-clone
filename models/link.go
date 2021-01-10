package models

import "time"

type Link struct {
	ID uint `json:"id"`
	UserID uint `json:"user_id"`
	User User
	Url string `gorm:"type:text" json:"url"`
	ShortUrl string `gorm:"type:varchar(255)" json:"short_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}