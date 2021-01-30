package models

import (
	config "bitly-clone/configs"
	"gorm.io/gorm"
	"time"
)

type Link struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Url       string    `gorm:"type:text" json:"url"`
	ShortUrl  string    `gorm:"type:varchar(255)" json:"short_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (l *Link) AfterFind(db *gorm.DB) (err error) {
	linkForUser(l)

	return nil
}

func (l *Link) AfterCreate(db *gorm.DB) (err error) {
	linkForUser(l)

	return nil
}

func linkForUser(l *Link) {
	if l.ShortUrl != "" {
		l.ShortUrl = config.BaseURL + "/" + l.ShortUrl
	}
}
