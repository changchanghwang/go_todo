package routes

import "github.com/gofiber/fiber/v2"

func SetUp(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})
}
