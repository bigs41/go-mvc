package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Web(app *fiber.App) {
	api := app
	api.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!--",
		})
	})

}
