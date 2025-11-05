package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome Home",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", home)

	router.Run("0.0.0.0:8081")
}
