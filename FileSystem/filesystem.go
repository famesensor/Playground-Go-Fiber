package FileSystem

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func main() {
	app := fiber.New()

	// Provide a minimal config
	app.Use(filesystem.New(filesystem.Config{
		Root: http.Dir("./assets"),
	}))

	// Or extend your config for customization
	app.Use(filesystem.New(filesystem.Config{
		Root:         http.Dir("./assets"),
		Browse:       true,
		Index:        "index.html",
		NotFoundFile: "404.html",
		MaxAge:       3600,
	}))
}
