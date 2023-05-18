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

	r.POST("/todos", func(c *gin.Context) {
		var newTodo Todo

		if err := c.BindJSON(&newTodo); err != nil {
			return
		}

		todos = append(todos, newTodo)
		c.IndentedJSON(http.StatusCreated, newTodo)
	})

	r.GET("/test", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "This is test api!!",
		})
	})

	r.Run()
}
