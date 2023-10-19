package main

import (
	"todo/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetUp(app)

	app.Listen(":3000")
}
