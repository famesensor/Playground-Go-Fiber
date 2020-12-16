package Favicon

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
)

func main() {
	app := fiber.New()

	// Provide a minimal config
	app.Use(favicon.New())

	// Or extend your config for customization
	app.Use(favicon.New(favicon.Config{
		File: "./favicon.ico",
	}))
}
