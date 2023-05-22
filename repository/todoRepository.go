package repository

import "todo-crud/models"

type TodoRepository interface {
	FindAll() []models.Todo
	Save(todo models.Todo)
	Update(todo models.Todo)
	Delete(todoId int)
}
