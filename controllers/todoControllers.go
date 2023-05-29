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

func (controllers *TodoController) GetAll(c *gin.Context) {
	todos := controllers.todoService.GetAll()

	response := response.ResponseBody{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   todos,
	}

	c.JSON(http.StatusOK, response)
}

func (controllers *TodoController) Create(c *gin.Context) {
	createTodoRequest := request.CreateTodoRequest{}
	err := c.ShouldBindJSON(&createTodoRequest)
	helper.LogAndPanicError(err)

	controllers.todoService.Create(createTodoRequest)

	response := response.ResponseBody{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   createTodoRequest,
	}

	c.JSON(http.StatusCreated, response)
}

func (controllers *TodoController) Update(c *gin.Context) {
	updateTodoRequest := request.UpdateTodoRequest{}
	err := c.ShouldBindJSON(&updateTodoRequest)
	helper.LogAndPanicError(err)

	todoId := c.Param("id")
	id, err := strconv.Atoi(todoId)
	helper.LogAndPanicError(err)
	println(id)
	// updateTodoRequest.Id = id

	controllers.todoService.Update(updateTodoRequest)

	response := response.ResponseBody{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   updateTodoRequest,
	}

	c.JSON(http.StatusOK, response)
}

func (controllers *TodoController) Delete(c *gin.Context) {
	todoId := c.Param("id")
	id, err := strconv.Atoi(todoId)
	helper.LogAndPanicError(err)

	controllers.todoService.Delete(id)

	response := response.ResponseBody{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   "Deleted!!",
	}

	c.JSON(http.StatusOK, response)
}

func (controllers *TodoController) TestController(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "This is test api!!",
	})
}
