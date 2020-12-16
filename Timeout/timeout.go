package Timeout

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

func main() {
	app := fiber.New()

	handler := func(ctx *fiber.Ctx) {
		ctx.SendString("Hello, World 👋!")
	}

	app.Get("/foo", timeout.New(handler, 5*time.Second))
}
