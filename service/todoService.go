package service

import (
	"todo-crud/request"
	"todo-crud/response"
)

type TodoService interface {
	GetAll() []response.TodoResponse
	GetById(todoId int) response.TodoResponse
	Create(todo request.CreateTodoRequest) response.TodoResponse
	Update(todoId int, todo request.UpdateTodoRequest) response.TodoResponse
	Delete(todoId int)
}
