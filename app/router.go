package app

import "github.com/gofiber/fiber/v2"

func InitRouter(app *fiber.App) {
	mainRoute(app)
}

func mainRoute(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", func(ctx *fiber.Ctx) error {
		err := ctx.JSON(fiber.Map{
			"code":    200,
			"message": "ok",
			"data":    nil,
		})
		return err
	})
	authRoute(api)
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

func featureRoute(app fiber.Router) {

}
