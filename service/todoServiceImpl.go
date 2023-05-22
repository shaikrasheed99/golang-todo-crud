package service

import (
	"todo-crud/models"
	"todo-crud/repository"
)

type TodoServiceImpl struct {
	todoRepository repository.TodoRepository
}

func NewTodoServiceImpl(todoRepository repository.TodoRepository) TodoService {
	return &TodoServiceImpl{todoRepository: todoRepository}
}

func (t *TodoServiceImpl) GetAll() []models.Todo {
	result := t.todoRepository.FindAll()

	return result
}

func (t *TodoServiceImpl) Create(todo models.Todo) {
	t.todoRepository.Save(todo)
}

func (t *TodoServiceImpl) Update(todo models.Todo) {
	t.todoRepository.Update(todo)
}

func (t *TodoServiceImpl) Delete(todoId int) {
	t.todoRepository.Delete(todoId)
}
