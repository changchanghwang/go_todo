package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()

	app.Get("/:name?", func(ctx *fiber.Ctx) error {
		if ctx.Params("name") != "" {
			return ctx.SendString("Hello " + ctx.Params("name"))
			// => Hello john
		}
		return ctx.SendString("Where is john?")
	})
	
	
	app.Listen(":3000")
}