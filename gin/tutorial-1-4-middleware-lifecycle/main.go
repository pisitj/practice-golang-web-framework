// https://github.com/gin-gonic/gin#custom-middleware

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// create a gin router with default middleware
	// logger & recovery (crash-free) middleware
	router := gin.Default()

	router.Use(TimestampLoggerMiddleware())

	router.GET("/", func(c *gin.Context) {
		fmt.Println("Request In-Process ..")
		c.JSON(http.StatusOK, gin.H{"message": "Hello World."})
	})

	// by default, it serves on :8080
	router.Run(":3000")
}

func TimestampLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		fmt.Printf("Request Start at %v \n", startTime)

		c.Next()

		endTime := time.Now()
		fmt.Printf("Request End at %v \n", endTime)
	}
}
