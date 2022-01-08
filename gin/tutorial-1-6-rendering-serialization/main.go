// Rendering & Serialization
// https://github.com/gin-gonic/gin#xml-json-yaml-and-protobuf-rendering

//https://pkg.go.dev/github.com/gin-gonic/gin#Context
// c.String
// c.JSON
// c.XML
// c.YAML
// c.ProtoBuf
// c.HTML

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	// create a gin router with default middleware
	// logger & recovery (crash-free) middleware
	router := gin.Default()

	// gin.H is a shortcut for map[string]interface{}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World.")
	})

	router.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello JSON"})
	})

	router.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "Hello XML"})
	})

	router.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "Hello YAML"})
	})

	// by default, it serves on :8080
	router.Run(":3000")
}
