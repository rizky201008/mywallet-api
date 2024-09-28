package test

import (
	"fmt"
	"github.com/rizky201008/mywallet-backend/model/domain"
	"github.com/rizky201008/mywallet-backend/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

var repo repository.TransactionRepository = repository.NewTransactionRepository()

func TestFindAll(t *testing.T) {
	result, err := repo.FindAll(db)

	fmt.Println(result)

	assert.Nil(t, err)
}

func insertExampleData() {
	data := domain.Transaction{
		UserID: 1,
		Desc:   nil,
		Amount: 3000.00,
	}
	db.Save(&data)
}

func truncateAllData() {
	query := db.Exec("TRUNCATE TABLE transactions")
	if query.Error != nil {
		panic(query.Error)
	}
}

func TestFindSuccess(t *testing.T) {
	defer truncateAllData()
	insertExampleData()
	result, err := repo.Find(db, 1)

	fmt.Println(result)

	assert.Nil(t, err)
}

func TestFindError(t *testing.T) {
	defer truncateAllData()
	insertExampleData()
	_, err := repo.Find(db, 2)

	assert.NotNil(t, err)
	fmt.Println(err)
}

func TestCreateSuccess(t *testing.T) {
	defer truncateAllData()
	transaction := domain.Transaction{
		UserID: 1,
		Desc:   nil,
		Amount: 3000.00,
	}

	_, err := repo.Create(db, transaction)

	assert.Nil(t, err)
	fmt.Println(transaction)
}

func TestCreateFailed(t *testing.T) {
	defer truncateAllData()
	transaction := domain.Transaction{
		UserID: 100,
		Desc:   nil,
		Amount: 3000.00,
	}

	err := db.Create(&transaction).Error

	assert.NotNil(t, err)
	fmt.Println(err)
}

func TestUpdateSuccess(t *testing.T) {
	defer truncateAllData()
	insertExampleData()
	transaction := domain.Transaction{
		UserID: 1,
		Desc:   nil,
		Amount: 32000.00,
	}
	_, err := repo.Update(db, transaction, 1)
	assert.Nil(t, err)
	fmt.Println(transaction)
}

func TestUpdateFailed(t *testing.T) {
	defer truncateAllData()
	insertExampleData()
	transaction := domain.Transaction{
		UserID: 13,
		Desc:   nil,
		Amount: 32000.00,
	}
	_, err := repo.Update(db, transaction, 1)
	assert.NotNil(t, err)
	fmt.Println(err)
}

func TestDeleteSuccess(t *testing.T) {
	defer truncateAllData()
	insertExampleData()
	err := repo.Delete(db, 1)

	assert.Nil(t, err)
}
