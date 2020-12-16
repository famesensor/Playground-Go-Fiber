package Cache

import (
	"time"

	"github.com/gofiber/fiber/v2"
	cache "github.com/gofiber/fiber/v2/middleware/cache"
)

func main() {
	app := fiber.New()

	// Initialize default config
	app.Use(cache.New())

	// Or extend your config for customization
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))
}
