package main

import (
	"github.com/rizky201008/mywallet-backend/app"
	"github.com/rizky201008/mywallet-backend/controller"
	"github.com/rizky201008/mywallet-backend/repository"
	"github.com/rizky201008/mywallet-backend/service"
)

func main() {
	viper := app.InitViper()
	db := app.InitDb(viper)
	transactionRepo := repository.NewTransactionRepository()
	transactionService := service.NewTransactionService(db, transactionRepo)
	transactionController := controller.NewTransactionController(transactionService)
	app.InitFiber(viper, transactionController)
}
