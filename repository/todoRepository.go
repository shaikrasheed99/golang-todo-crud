package repository

import "todo-crud/models"

type TodoRepository interface {
	FindAll() []models.Todo
	FindById(todoId int) models.Todo
	Save(todo models.Todo) models.Todo
	Update(todoId int, todo models.Todo)
	Delete(todoId int)
}
