package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizky201008/mywallet-backend/controller"
	"github.com/spf13/viper"
)

func InitFiber(viper *viper.Viper, transactionController controller.TransactionController) {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Redirect("/api")
	})

	InitRouter(app, transactionController)

	err := app.Listen(viper.GetString("app.APP_PORT"))

	if err != nil {
		panic(err)
	}
}
