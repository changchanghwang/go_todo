package todos

import (
	"encoding/json"
	"todo/services/todos/application"

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
		type request struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		}

		var r request
		err := ctx.BodyParser(&r)
		if err != nil {
			return err
		}

		application.Add(r.Title, r.Description)

		return ctx.SendString("ok")
	})
}
