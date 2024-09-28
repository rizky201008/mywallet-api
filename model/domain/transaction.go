package domain

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Amount float64 `gorm:"column:amount;type:double;default:0.0;"`
	UserID int     `gorm:"column:user_id;"`
	User   User    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Desc   *string `gorm:"column:desc;type:varchar(255);"`
}
