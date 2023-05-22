package service

import "todo-crud/models"

type TodoService interface {
	GetAll() []models.Todo
	Create(todo models.Todo)
	Update(todo models.Todo)
	Delete(todoId int)
}
