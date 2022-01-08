// Binding & Validation
// https://github.com/gin-gonic/gin#model-binding-and-validation

// https://pkg.go.dev/github.com/gin-gonic/gin#Context
// ShouldBind
// ShouldBindHeader
// ShouldBindQuery
// ShouldBindUri
// ShouldBindJSON
// ShouldBindXML
// ShouldBindYAML

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

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World."})
	})

	loginRoute := router.Group("/login")
	{
		loginRoute.POST("/json", loginJSON)
		loginRoute.POST("/form-data", loginForm)
	}

	// by default, it serves on :8080
	router.Run(":3000")
}

func loginJSON(c *gin.Context) {
	var jsonData LoginData

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON Data."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": jsonData.Username,
		"password": jsonData.Password,
	})
}

func loginForm(c *gin.Context) {
	var formData LoginData

	if err := c.ShouldBind(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Form Data."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": formData.Username,
		"password": formData.Password,
	})
}
