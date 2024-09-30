package repository

import (
	"github.com/rizky201008/mywallet-backend/model/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(db *gorm.DB, user domain.User) (domain.User, error)
	UpdateUser(db *gorm.DB, user domain.User) (domain.User, error)
	FindUserByUsername(db *gorm.DB, username string) (domain.User, error)
	FindUserById(db *gorm.DB, id int) (domain.User, error)
	DeleteUser(db *gorm.DB, user domain.User) error
}

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return UserRepositoryImpl{}
}

func (repository UserRepositoryImpl) CreateUser(db *gorm.DB, user domain.User) (domain.User, error) {
	created := db.Create(&user)
	return user, created.Error
}

func (repository UserRepositoryImpl) UpdateUser(db *gorm.DB, user domain.User) (domain.User, error) {
	updated := db.Where("username = ?", user.Username).Updates(&user)
	return user, updated.Error
}

func (repository UserRepositoryImpl) FindUserByUsername(db *gorm.DB, username string) (domain.User, error) {
	var user domain.User
	founded := db.Where("username = ?", username).First(&user)
	return user, founded.Error
}

func (repository UserRepositoryImpl) FindUserById(db *gorm.DB, id int) (domain.User, error) {
	var user domain.User
	founded := db.First(&user, id)
	return user, founded.Error
}

func (repository UserRepositoryImpl) DeleteUser(db *gorm.DB, user domain.User) error {
	return db.Delete(&domain.User{}, user.ID).Error
}
