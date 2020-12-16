package Proxy

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {
	app := fiber.New()
	// Forward to url
	app.Get("/gif", proxy.Forward("https://i.imgur.com/IWaBepg.gif"))

	// Make request within handler
	app.Get("/:id", func(c *fiber.Ctx) error {
		url := "https://i.imgur.com/" + c.Params("id") + ".gif"
		if err := proxy.Do(c, url); err != nil {
			return err
		}
		// Remove Server header from response
		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	})

	// Minimal round robin balancer
	app.Use(proxy.Balancer(proxy.Config{
		Servers: []string{
			"http://localhost:3001",
			"http://localhost:3002",
			"http://localhost:3003",
		},
	}))

	// Or extend your balancer for customization
	app.Use(proxy.Balancer(proxy.Config{
		Servers: []string{
			"http://localhost:3001",
			"http://localhost:3002",
			"http://localhost:3003",
		},
		ModifyRequest: func(c *fiber.Ctx) error {
			c.Request().Header.Add("X-Real-IP", c.IP())
			return nil
		},
		ModifyResponse: func(c *fiber.Ctx) error {
			c.Response().Header.Del(fiber.HeaderServer)
			return nil
		},
	}))
}
