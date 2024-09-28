package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);column:username;"`
	Password string `gorm:"type:varchar(16);column:password;"`
}
