// https://github.com/gofiber/fiber#grouping-routes-into-chains

package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/logger"
	"github.com/gofiber/fiber/middleware/recover"
)

func main() {
	app := fiber.New()

	// global-scope middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(ABCLoggerMiddleware)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello World."})
	}) // ABC

	v1Route := app.Group("/v1")
	{
		// group-scope middleware
		v1Route.Use(XYZLoggerMiddleware)

		v1Route.Get("/user", getUserList) // ABC XYZ
		v1Route.Post("/user", createUser) // ABC XYZ
	}

	v2Route := app.Group("/v2")
	{
		// endpoint-scope middleware
		v2Route.Get("/user", XYZLoggerMiddleware, getUserList) // ABC XYZ
		v2Route.Post("/user", createUser)                      // ABC
	}

	log.Fatal(app.Listen(":3000"))
}

func getUserList(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get List of Users."})
}

func createUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Create a User."})
}

func ABCLoggerMiddleware(c *fiber.Ctx) error {
	fmt.Println("ABC Logger Middleware.")
	return c.Next()
}

func XYZLoggerMiddleware(c *fiber.Ctx) error {
	fmt.Println("XYZ Logger Middleware.")
	return c.Next()
}
