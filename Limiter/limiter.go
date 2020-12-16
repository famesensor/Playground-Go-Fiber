package Limiter

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	app := fiber.New()

	// Default middleware config
	app.Use(limiter.New())

	// Or extend your config for customization
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:      20,
		Duration: 30 * time.Second,
		Key: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendFile("./toofast.html")
		},
		Store: myCustomStore{},
	}))
}
