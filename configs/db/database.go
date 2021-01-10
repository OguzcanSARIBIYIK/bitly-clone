package db

import (
	config "bitly-clone/configs"
	"bitly-clone/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var MyDB *gorm.DB

func init()  {
	connection()
	_ = MyDB.AutoMigrate(&models.User{})
	_ = MyDB.AutoMigrate(&models.Link{})
}

func connection(){
	dsn := config.DbUsername +":"+config.DbPassword+"@tcp("+config.DbIp+")/"+config.DbName+"?charset="+config.DbCharset+"&parseTime=True&loc=Local"
	var err error
	MyDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connection error..")
	}
}