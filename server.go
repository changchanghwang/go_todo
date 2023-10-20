package main

import (
	"todo/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetUp(app)

	err := app.Listen(":3000")
	if err != nil {
		print(err.Error())
		return
	}
}
