package routes

import (
	"todo/routes/todos"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	api := app.Group("/")

	todos.SetUp(app)

	api.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("pong")
	})
}
