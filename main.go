package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
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

	// Fiber instance
	app := fiber.New()
	// // Fiber config custom error
	// // Create a new fiber instance with custom config
	// app := fiber.New(fiber.Config{
	// 	// Override default error handler
	// 	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
	// 		// Statuscode defaults to 500
	// 		code := fiber.StatusInternalServerError

	// 		// Retreive the custom statuscode if it's an fiber.*Error
	// 		if e, ok := err.(*fiber.Error); ok {
	// 			code = e.Code
	// 		}

	// 		// Send custom error page
	// 		err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
	// 		if err != nil {
	// 			// In case the SendFile fails
	// 			return ctx.Status(500).SendString("Internal Server Error")
	// 		}

	// 		// Return from handler
	// 		return nil
	// 	},
	// })

	// ...
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

	// // Location
	// app.Post("/", func(c *fiber.Ctx) error {
	// 	// return c.Location("http://example.com")

	// 	return c.Location("/foo/bar")
	// })

	// Method
	app.Post("/method", func(c *fiber.Ctx) error {
		c.Method() // "POST"

		c.Method("GET")
		c.Method() // GET

		// ...
		return c.SendString(c.Method())
	})

	// MultipartForm && SaveFile
	app.Post("/multipartform", func(c *fiber.Ctx) error {
		// Parse the multipart form:
		if form, err := c.MultipartForm(); err == nil {
			// => *multipart.Form

			if token := form.Value["token"]; len(token) > 0 {
				// Get key value:
				fmt.Println(token[0])
			}

			// Get all files from "documents" key:
			files := form.File["documents"]
			// => []*multipart.FileHeader

			// Loop through files:
			for _, file := range files {
				fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
				// => "tutorial.pdf" 360641 "application/pdf"

				// Save the files to disk:
				if err := c.SaveFile(file, fmt.Sprintf("./%s", file.Filename)); err != nil {
					return err
				}
			}
		}

		return nil
	})

	// Params
	// GET http://localhost:3000/user/fenny
	app.Get("/user/:name", func(c *fiber.Ctx) error {
		c.Params("name") // "fenny"

		// ...
		return c.SendString(c.Params("name"))
	})

	app.Get("/v1/*/shop/*", func(c *fiber.Ctx) error {
		c.Params("*") // outputs the values of the first wildcard segment

		// ROUTE: /v1/*/shop/*
		// GET:   /v1/brand/4/shop/blue/xs
		c.Params("*1") // "brand/4"
		c.Params("*2") // "blue/xs"
		return c.SendString(c.Params("*4"))
	})

	// Path
	app.Get("/users", func(c *fiber.Ctx) error {
		c.Path() // "/users"

		c.Path("/john")
		c.Path() // "/john"

		// ...
		return c.SendString(c.Path())
	})

	// Query
	// GET http://localhost:3000/shoes?order=desc&brand=nike
	app.Get("/", func(c *fiber.Ctx) error {
		c.Query("order")         // "desc"
		c.Query("brand")         // "nike"
		c.Query("empty", "nike") // "nike"

		// ...
		return c.SendString(c.Query("order"))
	})

	//	QueryParser
	// Field names should start with an uppercase letter
	// curl -X POST "http://localhost:3000/?name=john&pass=doe&products=shoe,hat"
	type Person struct {
		Name     string   `query:"name"`
		Pass     string   `query:"pass"`
		Products []string `query:"products"`
	}

	app.Post("/", func(c *fiber.Ctx) error {
		p := new(Person)

		if err := c.QueryParser(p); err != nil {
			return err
		}

		log.Println(p.Name)     // john
		log.Println(p.Pass)     // doe
		log.Println(p.Products) // [shoe, hat]

		// ...
		return c.SendString(p.Name)
	})

	// Range
	// Range: bytes=500-700, 700-900
	app.Get("/", func(c *fiber.Ctx) error {
		b, _ := c.Range(1000)
		if b.Type == "bytes" {
			// for r := range r.Ranges {
			// 	fmt.Println(r)
			// 	// [500, 700]
			// }
			fmt.Print(b.Ranges)
		}
		return c.SendString("Range")
	})

	// Redirect
	app.Get("/coffee", func(c *fiber.Ctx) error {
		return c.Redirect("/teapot")
	})

	app.Get("/teapot", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusTeapot).SendString("🍵 short and stout 🍵")
	})

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Redirect("/foo/bar")
	// 	return c.Redirect("../login")
	// 	return c.Redirect("http://example.com")
	// 	return c.Redirect("http://example.com", 301)
	// })

	// Send && SendStream && SendString
	app.Get("/send", func(c *fiber.Ctx) error {
		return c.Send([]byte("Hello, World!")) // => "Hello, World!"
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
		// => "Hello, World!"

		// return c.SendStream(bytes.NewReader([]byte("Hello, World!")))
		// // => "Hello, World!"
	})

	// SendFile
	app.Get("/not-found", func(c *fiber.Ctx) error {
		return c.SendFile("./public/404.html")

		// // Disable compression
		// return c.SendFile("./static/index.html", false)
	})

	// SendStatus
	app.Get("/sendstatus", func(c *fiber.Ctx) error {
		return c.SendStatus(415)
		// => 415 "Unsupported Media Type"

		// c.SendString("Hello, World!")
		// return c.SendStatus(415)
		// // => 415 "Hello, World!"
	})

	// Status
	app.Get("/", func(c *fiber.Ctx) error {
		c.Status(200)
		return nil

		// return c.Status(400).Send([]byte("Bad Request"))

		// return c.Status(404).SendFile("./public/gopher.png")
	})

	// ErrorHandler
	// Catching Error
	app.Use(recover.New())
	app.Get("/errorhandler", func(c *fiber.Ctx) error {
		panic("This panic is catched by fiber")
	})

	app.Get("/newerror", func(c *fiber.Ctx) error {
		// 503 Service Unavailable
		return fiber.ErrServiceUnavailable

		// // 503 On vacation!
		// return fiber.NewError(fiber.StatusServiceUnavailable, "On vacation!")
	})

	// Listen
	app.Listen(":3000")
}
