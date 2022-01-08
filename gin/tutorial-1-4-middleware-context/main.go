// https://github.com/gin-gonic/gin#custom-middleware

// c.Get vs c.MustGet
// https://github.com/gin-gonic/gin/issues/117
// https://pkg.go.dev/github.com/gin-gonic/gin#Context.Get
// https://pkg.go.dev/github.com/gin-gonic/gin#Context.MustGet

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

	router.Use(TokenMiddleware())

	router.GET("/", func(c *gin.Context) {
		// get key-value from context
		val, exists := c.Get("token")
		if exists {
			token := val.(string)
			fmt.Printf("token from context >> %v \n", token)
		} else {
			fmt.Println("No key=token in context.")
		}

		c.JSON(http.StatusOK, gin.H{"message": "Hello World."})
	})

	// by default, it serves on :8080
	router.Run(":3000")
}

func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// set key-value in context
		c.Set("token", "12345")

		c.Next()
	}
}
