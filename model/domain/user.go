package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);column:username;unique;"`
	Password string `gorm:"type:text;column:password;"`
}
