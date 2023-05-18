package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
}

var todos = []Todo{
	{Id: 1, Description: "Sleeping", Priority: "high"},
	{Id: 3, Description: "Reading", Priority: "medium"},
	{Id: 2, Description: "Playing", Priority: "low"},
}

func main() {
	r := gin.Default()

	r.GET("/todos", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, todos)
	})

	r.GET("/test", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "This is test api!!",
		})
	})

	r.Run()
}
