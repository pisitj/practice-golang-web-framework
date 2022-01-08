// https://github.com/gin-gonic/gin#grouping-routes

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// create a gin router with default middleware
	// logger & recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World."})
	})

	v1Route := router.Group("/v1")
	{
		v1Route.GET("/user", getUserList)
		v1Route.POST("/user", createUser)
	}

	v2Route := router.Group("/v2")
	{
		v2Route.GET("/user", getUserList)
		v2Route.POST("/user", createUser)
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
