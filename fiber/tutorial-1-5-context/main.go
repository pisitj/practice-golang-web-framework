// https://docs.gofiber.io/api/ctx

// Accepts
// BaseURL
// Body
// BodyParser
// Cookies
// Request
// Response
// FormFile
// FormValue
// Get
// Hostname
// IP
// JSON
// OriginalURL
// Params
// Path
// Query
// QueryParser
// Set

package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/logger"
	"github.com/gofiber/fiber/middleware/recover"
)

type UserInfoData struct {
	Firstname string `query:"firstname" json:"firstname"`
	Lastname  string `query:"lastname" json:"lastname"`
}

type UserLoginData struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello World."})
	})

	app.Get("/user/:id", func(c *fiber.Ctx) error {
		fmt.Printf("hostname >> %v \n", c.Hostname())
		fmt.Printf("ip >> %v \n", c.IP())
		fmt.Printf("base url >> %v \n", c.BaseURL())
		fmt.Printf("original url >> %v \n", c.OriginalURL())
		fmt.Printf("request header - content-type >> %v \n", c.Get("Content-Type"))
		fmt.Printf("path parameter - id >> %v \n", c.Params("id"))

		return c.JSON(fiber.Map{"message": "Hello fiber.Ctx"})
	})

	// query parameter
	app.Get("/user", func(c *fiber.Ctx) error {
		fmt.Printf("query parameter - firstname >> %v \n", c.Query("firstname"))
		fmt.Printf("query parameter - lastname >> %v \n", c.Query("lastname"))

		userInfoData := new(UserInfoData)
		c.QueryParser(userInfoData)
		fmt.Printf("query parameter - firstname >> %v \n", userInfoData.Firstname)
		fmt.Printf("query parameter - lastname >> %v \n", userInfoData.Lastname)

		return c.JSON(&userInfoData)
	})

	// request body
	app.Post("/login", func(c *fiber.Ctx) error {
		fmt.Printf("request body >> %v \n", c.Body())

		userLoginData := new(UserLoginData)
		c.BodyParser(userLoginData)
		fmt.Printf("request body - username >> %v \n", userLoginData.Username)
		fmt.Printf("request body - password >> %v \n", userLoginData.Password)

		return c.JSON(&userLoginData)
	})

	log.Fatal(app.Listen(":3000"))
}
