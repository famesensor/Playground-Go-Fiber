package Basicauth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func auth() {
	app := fiber.New()

	// Provide a minimal config
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
	}))

	// Or extend your config for customization
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == "john" && pass == "doe" {
				return true
			}
			if user == "admin" && pass == "123456" {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.SendFile("./unauthorized.html")
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}))
}
