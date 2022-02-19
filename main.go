package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new fiber app
	app := fiber.New()
	port := os.Getenv("PORT")

	// Send a string back for GET calls to the endpoint "/"
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP!")
		return err
	})

	// Listen on PORT 3000
	app.Listen(port)
}
