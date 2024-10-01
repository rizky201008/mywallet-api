package app

import (
	"github.com/rizky201008/mywallet-backend/service"
)

var TransactionService service.TransactionService
var UserService service.UserService

func InitService() {
	TransactionService = service.NewTransactionService(Db, TransactionRepo)
	UserService = service.NewUserService(Db, UserRepo, Vipers)
}
