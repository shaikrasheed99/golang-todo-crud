package controllers

import (
	"net/http"
	"strconv"
	"todo-crud/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Todos)
}

func AddTodo(c *gin.Context) {
	var newTodo models.Todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	models.Todos = append(models.Todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func UpdateTodoById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	for index, todo := range models.Todos {
		if todo.Id == id {
			var newTodo models.Todo
			c.BindJSON(&newTodo)
			newTodo.Id = id
			models.Todos = append(models.Todos[:index], models.Todos[index+1:]...)
			models.Todos = append(models.Todos, newTodo)
			c.IndentedJSON(http.StatusCreated, newTodo)
			return
		}
	}

	c.IndentedJSON(http.StatusBadRequest, gin.H{
		"message": "Todo id is not found",
	})
}

func DeleteTodoById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for index, todo := range models.Todos {
		if todo.Id == id {
			models.Todos = append(models.Todos[:index], models.Todos[index+1:]...)
			c.IndentedJSON(http.StatusCreated, gin.H{
				"message": "Deleted!!",
			})
			return
		}
	}

	c.IndentedJSON(http.StatusBadRequest, gin.H{
		"message": "Todo id is not found",
	})
}

func TestController(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "This is test api!!",
	})
}
