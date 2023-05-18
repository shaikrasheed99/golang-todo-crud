package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "This is test api!!",
		})
	})

	r.Run()
}
