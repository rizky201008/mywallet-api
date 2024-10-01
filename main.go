package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rizky201008/mywallet-backend/app"
	"github.com/rizky201008/mywallet-backend/exception"
)

func init() {
	app.InitViper()
	app.InitRepository()
	app.InitDb()
	app.InitService()
	app.InitController()
}

func main() {
	viper := app.Vipers
	apps := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})

	apps.Use(recover.New())

	app.MainRouter(apps)

	err := apps.Listen(viper.GetString("app.APP_PORT"))

	if err != nil {
		panic(err)
	}
}
