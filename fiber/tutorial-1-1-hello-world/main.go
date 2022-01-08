// https://github.com/gofiber/fiber#json-response

package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/logger"
	"github.com/gofiber/fiber/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello World."})
	})

	log.Fatal(app.Listen(":3000"))
}
