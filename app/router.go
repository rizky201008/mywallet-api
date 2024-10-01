package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky201008/mywallet-backend/middleware"
)

func MainRouter(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Redirect("/api")
	})
	api := app.Group("/api")
	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"status":  "0",
			"message": "success",
			"data":    nil,
		})
	})
	authRoute(api)
	transactionRoute(api)
}

func authRoute(app fiber.Router) {
	userController := UserController
	auth := app.Group("/auth")
	auth.Post("/login", userController.Login)
	auth.Post("/register", userController.Register)
}

func transactionRoute(app fiber.Router) {
	transactionController := TransactionController
	transaction := app.Group("/transaction")
	transaction.Use(func(ctx *fiber.Ctx) error {
		return middleware.RequireAuth(ctx, Vipers, Db)
	})
	transaction.Get("/", transactionController.GetAllTransactions)
	transaction.Get("/:id", transactionController.GetTransaction)
	transaction.Post("/", transactionController.CreateTransaction)
	transaction.Put("/:id", transactionController.UpdateTransaction)
	transaction.Delete("/:id", transactionController.DeleteTransaction)
}
