package test

import (
	"fmt"
	"github.com/rizky201008/mywallet-backend/model/domain"
	"github.com/rizky201008/mywallet-backend/repository"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

var userRepository = repository.NewUserRepository()

func insertExampleUserData() {
	data := domain.User{
		Username: "joko",
		Password: "budi",
	}
	db.Save(&data)
}

func truncateAllUserData() {
	query := db.Exec("SET FOREIGN_KEY_CHECKS=0;") // ignore foreign_key checks
	if query.Error != nil {
		panic(query.Error)
	}
	query = db.Exec("TRUNCATE TABLE users;")
	if query.Error != nil {
		panic(query.Error)
	}
	query = db.Exec("SET FOREIGN_KEY_CHECKS=1;")
	if query.Error != nil {
		panic(query.Error)
	}
}

func TestGenerateHashedPassword(t *testing.T) {
	password := "Admin123#"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	assert.Nil(t, err)
	fmt.Println(string(hashedPassword))

	matchingPassword := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	assert.Nil(t, matchingPassword)
}

func TestCreateUserSuccess(t *testing.T) {
	defer truncateAllUserData()
	user := domain.User{
		Username: "jokowi",
		Password: "budi",
	}

	createdUser, err := userRepository.CreateUser(db, user)
	assert.Nil(t, err)
	fmt.Println(createdUser)
}

func TestCreateUserFailedUserExists(t *testing.T) {
	defer truncateAllUserData()
	insertExampleUserData()
	user := domain.User{
		Username: "jokowi",
		Password: "budi",
	}

	user.ID = 1 // inserting a new data with id 1

	_, err := userRepository.CreateUser(db, user)
	assert.NotNil(t, err)
	fmt.Println(err)
}

func TestFindUserById(t *testing.T) {
	defer truncateAllUserData()
	insertExampleUserData()
	user, err := userRepository.FindUserById(db, 1)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, uint(1))
}

func TestFindUserFailedNotFound(t *testing.T) {
	defer truncateAllUserData()
	insertExampleUserData()
	_, err := userRepository.FindUserById(db, 2)
	assert.NotNil(t, err)
}

func TestFindUserByUsername(t *testing.T) {
	defer truncateAllUserData()
	insertExampleUserData()
	user, err := userRepository.FindUserByUsername(db, "joko")
	assert.Nil(t, err)
	assert.Equal(t, user.Username, "joko")
}

func TestFindUserByUsernameFailedNotFound(t *testing.T) {
	defer truncateAllUserData()
	insertExampleUserData()
	_, err := userRepository.FindUserByUsername(db, "jokos")
	assert.NotNil(t, err)
}

func TestUpdateUserSuccess(t *testing.T) {
	defer truncateAllUserData()
	insertExampleUserData()
	user, err := userRepository.FindUserById(db, 1)
	assert.Nil(t, err)
	user.Username = "changed"
	updatedUser, err := userRepository.UpdateUser(db, user)
	assert.Nil(t, err)
	assert.Equal(t, user.Username, updatedUser.Username)
}

func TestDeleteUserSuccess(t *testing.T) {
	defer truncateAllUserData()
	insertExampleUserData()
	user, err := userRepository.FindUserById(db, 1)
	assert.Nil(t, err)
	err = userRepository.DeleteUser(db, user)
	assert.Nil(t, err)
}
