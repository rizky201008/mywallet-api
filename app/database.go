package app

import (
	"github.com/rizky201008/mywallet-backend/model/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() {
	viper := Vipers
	dsn := viper.GetString("db.USERNAME") + ":" + viper.GetString("db.PASSWORD") + "@tcp(" + viper.GetString("db.HOST") + ":" + viper.GetString("db.PORT") + ")/" + viper.GetString("db.NAME") + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	err = db.AutoMigrate(&domain.User{}, domain.Transaction{})
	if err != nil {
		panic(err.Error())
	}
	Db = db
}
