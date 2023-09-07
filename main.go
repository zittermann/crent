package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/pug/v2"
	// "net/http" // embedded system
)

func main() {
	// Create a new engine
	engine := pug.New("./views", ".pug")

	// Or from an embedded system
	// See github.com/gofiber/embed for examples
	// engine := pug.NewFileSystem(http.Dir("./views", ".pug"))

	// Pass the engine to the views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"msg": "Hello, World!",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
