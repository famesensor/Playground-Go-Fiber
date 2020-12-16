package Pprof

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func main() {
	app := fiber.New()

	// Default middleware
	app.Use(pprof.New())
}
