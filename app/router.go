package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky201008/mywallet-backend/controller"
)

func InitRouter(app *fiber.App, transactionController controller.TransactionController) {
	api := app.Group("/api")
	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"status":  "0",
			"message": "success",
			"data":    nil,
		})
	})
	authRoute(api)
	transactionRoute(api, transactionController)
}

func authRoute(app fiber.Router) {
	auth := app.Group("/auth")
	auth.Post("/login", func(ctx *fiber.Ctx) error {
		return ctx.SendString("You are logged in")
	})
	auth.Post("/logout", func(ctx *fiber.Ctx) error {
		return ctx.SendString("You are logged out")
	})
}

func transactionRoute(app fiber.Router, transactionController controller.TransactionController) {
	transaction := app.Group("/transaction")
	transaction.Get("/", transactionController.GetAllTransactions)
	transaction.Get("/:id", transactionController.GetTransaction)
	transaction.Post("/", transactionController.CreateTransaction)
	transaction.Put("/:id", transactionController.UpdateTransaction)
	transaction.Delete("/:id", transactionController.DeleteTransaction)
}
