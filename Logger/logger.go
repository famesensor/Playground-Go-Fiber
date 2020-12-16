package Logger

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New()

	// Logging Request ID
	// Default middleware config
	app.Use(logger.New())
	app.Use(requestid.New())

	// ​app​.​Use​(​logger​.​New​(logger.​Config​{
	// 	// For more options, see the Config section
	// 	Format​: `${pid} ${locals:requestid} ${status} - ${method} ${path}​\n​`,
	// }))

	// Changing TimeZone & TimeFormat
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/New_York",
	}))

	// Custom File Writer
	file, err := os.OpenFile("./123.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	app.Use(logger.New(logger.Config{
		Output: file,
	}))
}
