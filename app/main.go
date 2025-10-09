package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		response := fiber.Map{
			"message":   "My name is Susi Sandoval",
			"timestamp": time.Now().Unix(),
		}
		return c.JSON(response)
	})

	app.Listen(":3000")
}
