package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky201008/mywallet-backend/service"
)

type TransactionController interface {
	GetAllTransactions(ctx *fiber.Ctx) error
	GetTransaction(ctx *fiber.Ctx) error
	CreateTransaction(ctx *fiber.Ctx) error
	UpdateTransaction(ctx *fiber.Ctx) error
	DeleteTransaction(ctx *fiber.Ctx) error
}

type TransactionControllerImpl struct {
	Service service.TransactionService
}

func NewTransactionController(service service.TransactionService) *TransactionControllerImpl {
	return &TransactionControllerImpl{
		Service: service,
	}
}

func (controller TransactionControllerImpl) GetAllTransactions(ctx *fiber.Ctx) error {
	response := controller.Service.GetAllTransaction()

	return ctx.JSON(fiber.Map{
		"status":  "0",
		"message": "success",
		"data":    response,
	})
}

func (controller TransactionControllerImpl) GetTransaction(ctx *fiber.Ctx) error {
	response := controller.Service.GetTransaction(ctx)
	return ctx.JSON(fiber.Map{
		"status":  "0",
		"message": "success",
		"data":    response,
	})
}

func (controller TransactionControllerImpl) CreateTransaction(ctx *fiber.Ctx) error {
	response := controller.Service.CreateTransaction(ctx)
	return ctx.JSON(fiber.Map{
		"status":  "0",
		"message": "success",
		"data":    response,
	})
}

func (controller TransactionControllerImpl) UpdateTransaction(ctx *fiber.Ctx) error {
	response := controller.Service.UpdateTransaction(ctx)
	return ctx.JSON(fiber.Map{
		"status":  "0",
		"message": "success",
		"data":    response,
	})
}

func (controller TransactionControllerImpl) DeleteTransaction(ctx *fiber.Ctx) error {
	controller.Service.DeleteTransaction(ctx)
	return ctx.JSON(fiber.Map{
		"status":  "0",
		"message": "success",
		"data":    nil,
	})
}
