package service

import (
	"todo-crud/request"
	"todo-crud/response"
)

type TodoService interface {
	GetAll() []response.TodoResponse
	Create(todo request.CreateTodoRequest)
	Update(todo request.UpdateTodoRequest)
	Delete(todoId int)
}
