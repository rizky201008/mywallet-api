package repository

import (
	"github.com/rizky201008/mywallet-backend/model/domain"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindAll(db *gorm.DB) ([]domain.Transaction, error)
	Find(db *gorm.DB, id int) (domain.Transaction, error)
	Create(db *gorm.DB, transaction domain.Transaction) (domain.Transaction, error)
	Update(db *gorm.DB, transaction domain.Transaction, id int) (domain.Transaction, error)
	Delete(db *gorm.DB, id int) error
}

type TransactionRepositoryImpl struct{}

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (TransactionRepositoryImpl) FindAll(db *gorm.DB) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	find := db.Find(&transactions)

	return transactions, find.Error
}

func (TransactionRepositoryImpl) Find(db *gorm.DB, id int) (domain.Transaction, error) {
	var transaction domain.Transaction
	find := db.First(&transaction, id)
	return transaction, find.Error
}

func (TransactionRepositoryImpl) Create(db *gorm.DB, transaction domain.Transaction) (domain.Transaction, error) {
	create := db.Create(&transaction)
	return transaction, create.Error
}

func (TransactionRepositoryImpl) Update(db *gorm.DB, transaction domain.Transaction, id int) (domain.Transaction, error) {
	update := db.Where("id = ?", id).Updates(&transaction)
	return transaction, update.Error
}

func (TransactionRepositoryImpl) Delete(db *gorm.DB, id int) error {
	return db.Delete(&domain.Transaction{}, id).Error
}
