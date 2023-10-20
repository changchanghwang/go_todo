package todos

import (
	"encoding/json"
	"strconv"
	todoService "todo/services/todos/application"
	"todo/services/todos/dto"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	api := app.Group("/todos")

	api.Get("/", func(ctx *fiber.Ctx) error {
		todos := todoService.List()
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

		todoService.Add(todoCreateDto)

		return ctx.SendString("ok")
	})

	api.Patch("/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		var todoUpdateDto dto.TodoUpdateDto
		err := ctx.BodyParser(&todoUpdateDto)
		if err != nil {
			return err
		}
		numId, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		todoService.Update(numId, todoUpdateDto)

		return ctx.SendString("ok")
	})
}
