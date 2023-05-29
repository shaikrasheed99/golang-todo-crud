package routes

import (
	"todo-crud/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(todoController *controllers.TodoController) *gin.Engine {
	r := gin.Default()

	r.GET("/todos", todoController.GetAll)
	r.GET("/todos/:id", todoController.GetById)
	r.POST("/todos", todoController.Create)
	r.PUT("/todos/:id", todoController.Update)
	r.DELETE("/todos/:id", todoController.Delete)

	r.GET("/test", todoController.TestController)

	return r
}
