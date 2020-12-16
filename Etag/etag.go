package Etag

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
)

func main() {
	app := fiber.New()

	// Default middleware config
	app.Use(etag.New())

	// Get / receives Etag: "13-1831710635" in response header
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
