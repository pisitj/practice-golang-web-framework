// https://github.com/gofiber/fiber#grouping-routes-into-chains

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

	v1Route := app.Group("/v1")
	{
		v1Route.Get("/user", getUserList)
		v1Route.Post("/user", createUser)
	}

	v2Route := app.Group("/v2")
	{
		v2Route.Get("/user", getUserList)
		v2Route.Post("/user", createUser)
	}

	log.Fatal(app.Listen(":3000"))
}

func getUserList(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get List of Users."})
}

func createUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Create a User."})
}
