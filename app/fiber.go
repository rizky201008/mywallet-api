package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func InitFiber(viper *viper.Viper) {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Redirect("/api")
	})

	InitRouter(app)

	err := app.Listen(viper.GetString("app.APP_PORT"))

	if err != nil {
		panic(err)
	}
}
