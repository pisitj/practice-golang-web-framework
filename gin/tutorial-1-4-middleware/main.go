// https://gin-gonic.com/docs/examples/custom-middleware/

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// create a gin router with default middleware
	// logger & recovery (crash-free) middleware
	router := gin.Default()

	// global-scope middleware
	router.Use(ABCLoggerMiddleware())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World."})
	}) // ABC

	v1Route := router.Group("/v1")
	{
		// group-scope middleware
		v1Route.Use(XYZLoggerMiddleware())

		v1Route.GET("/user", getUserList) // ABC XYZ
		v1Route.POST("/user", createUser) // ABC XYZ
	}

	v2Route := router.Group("/v2")
	{
		// endpoint-scope middleware
		v2Route.GET("/user", XYZLoggerMiddleware(), getUserList) // ABC XYZ
		v2Route.POST("/user", createUser)                        // ABC
	}

	// by default, it serves on :8080
	router.Run(":3000")
}

func getUserList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get List of Users."})
}

func createUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create a User."})
}

func ABCLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("ABC Logger Middleware")
		c.Next()
	}
}

func XYZLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("XYZ Logger Middleware")
		c.Next()
	}
}
