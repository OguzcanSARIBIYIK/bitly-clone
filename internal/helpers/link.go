package helpers

import (
	"bitly-clone/configs/db"
	"bitly-clone/models"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func CreateLink() string {

	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	var user models.User
	db.MyDB.Model(&models.User{}).Where("token = ?", string(b)).First(&user)

	if user.ID == 0 {
		return string(b)
	}

	return CreateLink()
}
