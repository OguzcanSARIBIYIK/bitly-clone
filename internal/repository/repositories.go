package repository

import (
	"bitly-clone/configs/db"

	"gorm.io/gorm"
)

type Repositories struct {
	DB *gorm.DB
}

var Repo *Repositories

func Init() {
	Repo = &Repositories{
		DB: db.MyDB,
	}
}

func Get() *Repositories {
	return Repo
}
