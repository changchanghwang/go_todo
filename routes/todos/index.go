package todos

import (
	"encoding/json"
	"todo/services/todos/application"
	"todo/services/todos/dto"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	api := app.Group("/todos")

	api.Get("/", func(ctx *fiber.Ctx) error {
		todos := application.List()
		t, err := json.Marshal(todos)
		if err != nil {
			return err
		}
		return ctx.SendString(string(t))
	})

	api.Post("/", func(ctx *fiber.Ctx) error {

		var todoCreateDto dto.TodoCreateDto
		err := ctx.BodyParser(&todoCreateDto)
		if err != nil {
			return err
		}

		application.Add(todoCreateDto)

		return ctx.SendString("ok")
	})
}
