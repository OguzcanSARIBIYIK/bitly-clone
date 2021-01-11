package models

import "time"

type User struct {
	ID        uint      `json:"id"`
	Username  string    `gorm:"type:varchar(250)" json:"username"`
	Password  string    `json:"password"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
