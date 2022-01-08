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

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World."})
	})

	// query parameter
	// https://github.com/gin-gonic/gin#querystring-parameters
	router.GET("/welcome", func(c *gin.Context) {
		user := c.DefaultQuery("user", "Guest")
		firstname := c.Query("firstname") // shortcut for c.Request.URL.Query().Get("firstname")
		lastname := c.Query("lastname")   // shortcut for c.Request.URL.Query().Get("lastname")

		c.JSON(http.StatusOK, gin.H{
			"message":   "Welcome",
			"user":      user,
			"firstname": firstname,
			"lastname":  lastname,
		})
	})

	// path parameter
	// https://github.com/gin-gonic/gin#querystring-parameters
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		userMessage := fmt.Sprintf("Hello, %v", name)
		c.JSON(http.StatusOK, gin.H{
			"message": userMessage,
		})
	})

	// form-data
	// https://github.com/gin-gonic/gin#another-example-query--post-form
	router.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	// gin.Context
	// https://pkg.go.dev/github.com/gin-gonic/gin#Context

	// type Context struct {
	// 		Request *http.Request
	// 		Writer  ResponseWriter

	// 		Params
	// 		Keys map[string]interface{}
	// 		Errors errorMsgs
	// 		Accepted []string
	// }

	router.POST("/", func(c *gin.Context) {
		fmt.Sprintf("host >> %v \n", c.Request.Host)
		fmt.Sprintf("url >> %v \n", c.Request.URL)
		fmt.Sprintf("request header >> %v \n", c.Request.Header)
		fmt.Sprintf("request body >> %v \n", c.Request.Body)

		c.JSON(http.StatusOK, gin.H{
			"message": "Hiiii - log request detail in terminal.",
		})
	})

	// by default, it serves on :8080
	router.Run(":3000")
}
