package Session

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func main() {
	app := fiber.New()

	// Default middleware config
	store := session.New()

	// This panic will be catch by the middleware
	app.Get("/", func(c *fiber.Ctx) error {
		// get session from storage
		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}

		// save session
		defer sess.Save()

		// Get value
		name := sess.Get("name")

		// Set key/value
		sess.Set("name", "john")

		// Delete key
		sess.Delete("name")

		// Destry session
		if err := sess.Destroy(); err != nil {
			panic(err)
		}

		return fmt.Fprintf(ctx, "Welcome %v", name)
	})
}
