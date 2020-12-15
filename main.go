package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Bodyparser
type Person struct {
	Name string `json:"name" xml:"name" form:"name"`
	Pass string `json:"pass" xml:"pass" form:"pass"`
}

// Cookie
type Cookie struct {
	Name     string    `json:"name"`
	Value    string    `json:"value"`
	Path     string    `json:"path"`
	Domain   string    `json:"domain"`
	MaxAge   int       `json:"max_age"`
	Expires  time.Time `json:"expires"`
	Secure   bool      `json:"secure"`
	HTTPOnly bool      `json:"http_only"`
	SameSite string    `json:"same_site"`
}

// JSON
type SomeStruct struct {
	Name string
	Age  uint8
}

func main() {
	// // Mount
	// micro := fiber.New()

	// micro.Get("/hook", func(c *fiber.Ctx) error {
	// 	return c.SendString("I'm Hook")
	// })

	app := fiber.New()

	// Basic route
	app.Get("/", func(c *fiber.Ctx) error {
		// fmt.Println(c.BaseURL())              // => http://google.com
		// fmt.Println(c.Get("X-Custom-Header")) // => hi

		return c.SendString("hello, World!")
	})

	// Basic Error
	app.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(404, "Not Found")
	})

	app.Post("/api/register", func(c *fiber.Ctx) error {
		return c.SendString("I'm a POST request!")
	})

	// // Basic middleware : Use
	// app.Use(func(c *fiber.Ctx) error {
	// 	return c.SendString("Match Any thing!!")
	// })

	// app.Use("/api", func(c *fiber.Ctx) error {
	// 	return c.SendString("Match request starting with /api")
	// })

	// // Attach multiple handles
	// app.Use("/api", func(c *fiber.Ctx) error {
	// 	c.Set("X-Custom-Header")
	// 	c.Next()
	// }, func(c *fiber.Ctx) error {
	// 	return c.Next()
	// })

	// // Mount
	// app.Mount("/student", micro)

	// // Group
	// api := app.Group("/api", handler) // /api

	// v1 := api.Group("/v1", handler) // /api/v1
	// v1.Get("/list", handler)        // /api/v1/list
	// v1.Get("/user", handler)        // /api/v1/user

	// v2 := api.Group("/v2", handler) // /api/v2
	// v2.Get("/list", handler)        // /api/v2/list
	// v2.Get("/user", handler)        // /api/v2/user

	// // Stack
	// data, _ := json.MarshalIndent(app.Stack(), "", "")
	// fmt.Print(string(data))

	// // Test
	// // http.Request
	// req := httptest.NewRequest("GET", "http://google.com", nil)
	// req.Header.Set("X-Custom-Header", "hi")

	// // http.Response
	// resp, _ := app.Test(req)

	// // Do something with results:
	// if resp.StatusCode == 200 {
	// 	body, _ := ioutil.ReadAll(resp.Body)
	// 	fmt.Println(string(body)) // => Hello, World!
	// }

	// Accepts
	app.Get("/accept", func(c *fiber.Ctx) error {
		c.Accepts("html")             // "html"
		c.Accepts("text/html")        // "text/html"
		c.Accepts("json", "text")     // "json"
		c.Accepts("application/json") // "application/json"
		c.Accepts("image/png")        // ""
		c.Accepts("png")              // ""
		//...
		c.AcceptsCharsets("utf-16", "iso-8859-1")
		// "iso-8859-1"

		c.AcceptsEncodings("compress", "br")
		// "compress"

		c.AcceptsLanguages("pt", "nl", "ru")
		// "nl"
		// ...
		return nil
	})

	// Append
	app.Get("/append", func(c *fiber.Ctx) error {
		c.Append("Link", "google.com") // => google.com

		c.Append("Link", "facebook.com") // => google.com, facebook.com

		c.Append("Link", "youtube.com") // => google.com, facebook.com, youtube.comå
		return nil
	})

	// Attachement
	app.Get("/attachement", func(c *fiber.Ctx) error {
		c.Attachment()
		// => Content-Disposition: attachment

		c.Attachment("./upload/images/logo.png")
		// => Content-Disposition: attachment; filename="logo.png"
		// => Content-Type: image/png
		return nil
	})

	// BaseURL
	app.Get("/baseurl", func(c *fiber.Ctx) error {
		fmt.Print(c.BaseURL())

		return nil
	})

	// Body
	app.Post("/body", func(c *fiber.Ctx) error {
		fmt.Print(c.Body())

		return c.Send(c.Body())
	})

	// Bodyparser
	app.Post("/bodyparser", func(c *fiber.Ctx) error {
		p := new(Person)

		if err := c.BodyParser(p); err != nil {
			return err
		}

		log.Print(p.Name)
		log.Print(p.Pass)

		return nil
	})

	// Clear Cookie
	app.Get("/clear", func(c *fiber.Ctx) error {
		// Clears all cookies:
		c.ClearCookie()

		// Expire specific cookie by name:
		c.ClearCookie("user")

		// Expire multiple cookies by names:
		c.ClearCookie("token", "session", "track_id", "version")
		// ...
		return nil
	})

	app.Get("/set", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    "randomvalue",
			Expires:  time.Now().Add(24 * time.Hour),
			HTTPOnly: true,
			SameSite: "lax",
		})

		return nil
	})

	app.Get("/delete", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name: "token",
			// Set expiry date to the past
			Expires:  time.Now().Add(-(time.Hour * 2)),
			HTTPOnly: true,
			SameSite: "lax",
		})

		// ...
		return nil
	})

	// Cookie
	app.Get("/cookie", func(c *fiber.Ctx) error {
		// Create cookie
		cookie := new(fiber.Cookie)
		cookie.Name = "john"
		cookie.Value = "doe"
		cookie.Expires = time.Now().Add(24 * time.Hour)

		// Set cookie
		c.Cookie(cookie)

		return nil
	})

	// Cookies
	app.Get("/cookies", func(c *fiber.Ctx) error {
		// Get cookie by key:
		log.Print(c.Cookies("name"))
		c.Cookies("name") // "john"
		log.Print(c.Cookies("name"), c.Cookies("empty"))
		c.Cookies("empty", "Fame") // "Fame"
		log.Print(c.Cookies("empty"))
		// ...
		return nil
	})

	// Download
	app.Get("/download", func(c *fiber.Ctx) error {
		return c.Download("./files/report-12345.pdf")
		// => Download report-12345.pdf

		// return c.Download("./files/report-12345.pdf", "report.pdf")
		// // => Download report.pdf
	})

	// Request
	app.Get("/request", func(c *fiber.Ctx) error {
		log.Print(c.Request().Header.Method())

		return nil
	})

	// Response
	app.Get("/response", func(c *fiber.Ctx) error {
		c.Response()
		// => "Hello, World!"
		return nil
	})

	// Format
	app.Get("/format", func(c *fiber.Ctx) error {
		// Accept: text/plain
		c.Format("Hello, World!")
		// => Hello, World!

		// Accept: text/html
		c.Format("Hello, World!")
		// => <p>Hello, World!</p>

		// Accept: application/json
		c.Format("Hello, World!")
		// => "Hello, World!"
		// ..
		return nil
	})

	// Formfile
	app.Post("/formfile", func(c *fiber.Ctx) error {
		// Get first file from form field "document":
		file, err := c.FormFile("document")
		defer fmt.Print(err)
		// Save file to root directory:
		return c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
	})

	// FormValue
	app.Post("/formvalue", func(c *fiber.Ctx) error {
		// Get first value from form field "name":
		name := c.FormValue("name")
		pass := c.FormValue("pass")
		log.Print(name, pass)
		// => "john" or "" if not exist

		// ..
		return nil
	})

	// Get
	app.Get("/get", func(c *fiber.Ctx) error {
		fmt.Print(c.Get("Content-Type"))      // "text/plain"
		fmt.Print(c.Get("CoNtEnT-TypE"))      // "text/plain"
		fmt.Print(c.Get("something", "john")) // "john"
		// ..
		return nil
	})

	// JSON
	app.Get("/json", func(c *fiber.Ctx) error {
		// Create data struct:
		// 1
		data := SomeStruct{
			Name: "Grame",
			Age:  20,
		}
		// 2
		// data := new(SomeStruct)
		// data.Name = "Grame"
		// data.Age = 20
		// 3
		// var data SomeStruct
		// 4
		// temp := &SomeStruct{
		// 	Name: "fame",
		// 	Age:  20,
		// }

		// log.Panicln(temp)
		return c.JSON(data)
		// => Content-Type: application/json
		// => "{"Name": "Grame", "Age": 20}"

		// return c.JSON(fiber.Map{
		// 	"name": "Grame",
		// 	"age":  20,
		// })
		// => Content-Type: application/json
		// => "{"name": "Grame", "age": 20}"
	})

	// JSONP
	app.Get("/jsonp", func(c *fiber.Ctx) error {
		data := SomeStruct{
			Name: "Grame",
			Age:  20,
		}

		return c.JSONP(data) // => callback({Name: "Grame", Age: 20})
		// return c.JSONP(data, "customFunc") // => customFunc({Name: "Grame", Age: 20})
	})

	// // Link
	// app.Get("/link", func(c *fiber.Ctx) error {
	// 	c.Link(
	// 		"http://api.example.com/users?page=2", "next",
	// 		"http://api.example.com/users?page=5", "last",
	// 	)
	// 	// Link: <http://api.example.com/users?page=2>; rel="next",
	// 	//       <http://api.example.com/users?page=5>; rel="last"

	// 	// ...
	// 	return nil
	// })

	// Local
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("user", "admin")
		return c.Next()
	})

	app.Get("/admin", func(c *fiber.Ctx) error {
		if c.Locals("user") == "admin" {
			return c.Status(200).SendString("Welcome, admin!")
		}
		return c.SendStatus(403)

	})
	// Listen
	app.Listen(":3000")
}
