package controllers

import (
	"net/http"
	"strconv"
	"todo-crud/helper"
	"todo-crud/request"
	"todo-crud/response"
	"todo-crud/service"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoService service.TodoService
}

func NewTodoController(todoService service.TodoService) *TodoController {
	return &TodoController{todoService: todoService}
}

func (tc *TodoController) GetAll(c *gin.Context) {
	todos := tc.todoService.GetAll()

	response := response.ResponseBody{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   todos,
	}

	c.JSON(http.StatusOK, response)
}

func (tc *TodoController) GetById(c *gin.Context) {
	requestId := c.Param("id")

	todoId, err := strconv.Atoi(requestId)
	helper.LogAndPanicError(err)

	todo := tc.todoService.GetById(todoId)

	response := response.ResponseBody{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   todo,
	}

	c.IndentedJSON(http.StatusOK, response)
}

func (tc *TodoController) Create(c *gin.Context) {
	createTodoRequest := request.CreateTodoRequest{}
	err := c.ShouldBindJSON(&createTodoRequest)
	helper.LogAndPanicError(err)

	createdTodo := tc.todoService.Create(createTodoRequest)

	response := response.ResponseBody{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   createdTodo,
	}

	c.JSON(http.StatusCreated, response)
}

func (tc *TodoController) Update(c *gin.Context) {
	updateTodoRequest := request.UpdateTodoRequest{}
	err := c.ShouldBindJSON(&updateTodoRequest)
	helper.LogAndPanicError(err)

	requestId := c.Param("id")
	todoId, err := strconv.Atoi(requestId)
	helper.LogAndPanicError(err)

	updatedTodo := tc.todoService.Update(todoId, updateTodoRequest)

	response := response.ResponseBody{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   updatedTodo,
	}

	c.JSON(http.StatusOK, response)
}

func (tc *TodoController) Delete(c *gin.Context) {
	requestId := c.Param("id")
	todoId, err := strconv.Atoi(requestId)
	helper.LogAndPanicError(err)

	tc.todoService.Delete(todoId)

	response := response.ResponseBody{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   "Deleted!!",
	}

	c.JSON(http.StatusOK, response)
}

func (tc *TodoController) TestController(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "This is test api!!",
	})
}
