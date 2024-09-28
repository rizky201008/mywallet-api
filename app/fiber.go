package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rizky201008/mywallet-backend/controller"
	"github.com/rizky201008/mywallet-backend/exception"
	"github.com/spf13/viper"
)

func InitFiber(viper *viper.Viper, transactionController controller.TransactionController) {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})

	app.Use(recover.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Redirect("/api")
	})

	InitRouter(app, transactionController)

	err := app.Listen(viper.GetString("app.APP_PORT"))

	if err != nil {
		panic(err)
	}
}
