// https://github.com/gin-gonic/gin#using-get-post-put-patch-delete-and-options

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

	// by default, it serves on :8080
	router.Run(":3000")
}
