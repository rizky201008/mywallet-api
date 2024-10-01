package app

import "github.com/rizky201008/mywallet-backend/repository"

var UserRepo repository.UserRepository
var TransactionRepo repository.TransactionRepository

func InitRepository() {
	UserRepo = repository.NewUserRepository()
	TransactionRepo = repository.NewTransactionRepository()
}
