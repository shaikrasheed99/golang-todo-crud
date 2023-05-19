package routes

import (
	"todo-crud/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/todos", controllers.GetTodos)
	r.POST("/todos", controllers.AddTodo)
	r.PUT("/todos/:id", controllers.UpdateTodoById)
	r.DELETE("/todos/:id", controllers.DeleteTodoById)
	r.GET("/test", controllers.TestController)

	return r
}
