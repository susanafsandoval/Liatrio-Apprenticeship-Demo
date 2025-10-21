package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {

	/*
		app := fiber.New()

		app.Get("/", func(c *fiber.Ctx) error {
			response := fiber.Map{
				"message":   "My name is Susi Sandoval",
				"timestamp": time.Now().Unix(),
			}
			return c.JSON(response)
		})

		app.Listen(":3000")

	*/

	// Added dynamic port binding to align with Docker's PORT environment variable

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":   "My name is Susi Sandoval",
			"timestamp": time.Now().Unix(),
			"message1":  "Demo is Complete!",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app.Listen(":" + port)

}
