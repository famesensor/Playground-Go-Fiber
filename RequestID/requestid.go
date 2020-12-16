package RequestID

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New()

	// Default middleware config
	app.Use(requestid.New())

	// Or extend your config for customization
	app.Use(requestid.New(requestid.Config{
		Header: "X-Custom-Header",
		Generator: func() string {
			return "static-id"
		},
	}))
}
