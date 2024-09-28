package app

import (
	"github.com/rizky201008/mywallet-backend/model/domain"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb(viper *viper.Viper) *gorm.DB {
	dsn := viper.GetString("db.USERNAME") + ":" + viper.GetString("db.PASSWORD") + "@tcp(" + viper.GetString("db.HOST") + ":" + viper.GetString("db.PORT") + ")/" + viper.GetString("db.NAME") + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	err = db.AutoMigrate(&domain.User{}, domain.Transaction{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
