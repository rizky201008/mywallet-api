package app

import "github.com/rizky201008/mywallet-backend/controller"

var UserController controller.UserController
var TransactionController controller.TransactionController

func InitController() {
	UserController = controller.NewUserController(UserService)
	TransactionController = controller.NewTransactionController(TransactionService)
}
