package Recover

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	// Default middleware config
	app.Use(recover.New())

	// This panic will be catch by the middleware
	app.Get("/", func(c *fiber.Ctx) error {
		panic("I'm an error")
	})
}
