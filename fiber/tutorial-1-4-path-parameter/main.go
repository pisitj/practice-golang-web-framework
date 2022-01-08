// https://github.com/gofiber/fiber#-basic-routing

// Order of Router is very Matter !

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
		return c.JSON(fiber.Map{
			"message": "Hello World.",
		})
	})

	v1Route := app.Group("/v1")
	{
		v1Route.Get("/:name/:age", nameAgeHandler)
		v1Route.Get("/:filename.:extension", filenameExtensionHandler)
		v1Route.Get("/flights/:from-:to", flightsFromToHandler)
		v1Route.Get("/api/*", apiWildcardHandler)

		// GET v1/john/75             > /:name/:age >               {"name": "john", "age": "75"}
		// GET v1/dictionary.txt      > /:filename.:extension >     {"filename": "dictionary", "extension": "txt"}
		// GET v1/flights/LAX-SFO     > /:name/:age >               {"name": "flights", "age": "LAX-SFO"}
		// GET v1/api/register        > /:name/:age >               {"name": "api", "age": "register"}
	}

	v2Route := app.Group("/v2")
	{
		v2Route.Get("/:filename.:extension", filenameExtensionHandler)
		v2Route.Get("/api/*", apiWildcardHandler)
		v2Route.Get("/flights/:from-:to", flightsFromToHandler)
		v2Route.Get("/:name/:age", nameAgeHandler)

		// GET v2/dictionary.txt      > /:filename.:extension >     {"filename": "dictionary", "extension": "txt"}
		// GET v2/api/register        > /api/* >                    {"apiPath": "register"}
		// GET v2/flights/LAX-SFO     > /flights/:from-:to >        {"from": "LAX", "to": "SFO"}
		// GET v2/john/75             > /:name/:age >               {"name": "john", "age": "75"}

	}

	log.Fatal(app.Listen(":3000"))
}

// GET /:name/:age
func nameAgeHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"name": c.Params("name"),
		"age":  c.Params("age"),
	})
}

// GET /:filename.:extension
func filenameExtensionHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"filename":  c.Params("filename"),
		"extension": c.Params("extension"),
	})
}

// GET /flights/:from-:to
func flightsFromToHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"from": c.Params("from"),
		"to":   c.Params("to"),
	})
}

// GET /api/*
func apiWildcardHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"apiPath": c.Params("*"),
	})
}
