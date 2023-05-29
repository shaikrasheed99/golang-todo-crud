package service

import (
	"todo-crud/helper"
	"todo-crud/models"
	"todo-crud/repository"
	"todo-crud/request"
	"todo-crud/response"

	"github.com/go-playground/validator/v10"
)

type TodoServiceImpl struct {
	TodoRepository repository.TodoRepository
	Validate       *validator.Validate
}

func NewTodoServiceImpl(todoRepository repository.TodoRepository, validate *validator.Validate) TodoService {
	return &TodoServiceImpl{
		TodoRepository: todoRepository,
		Validate:       validate,
	}
}

func (t *TodoServiceImpl) GetAll() []response.TodoResponse {
	result := t.TodoRepository.FindAll()

	var todos []response.TodoResponse

	for _, value := range result {
		todo := response.TodoResponse{
			Id:          value.Id,
			Description: value.Description,
			Priority:    value.Priority,
		}

		todos = append(todos, todo)
	}

	return todos
}

func (t *TodoServiceImpl) Create(todo request.CreateTodoRequest) {
	validationError := t.Validate.Struct(todo)
	helper.LogAndPanicError(validationError)

	todoModel := models.Todo{
		Description: todo.Description,
		Priority:    todo.Priority,
	}

	t.TodoRepository.Save(todoModel)
}

func (t *TodoServiceImpl) Update(todo request.UpdateTodoRequest) {
	validationError := t.Validate.Struct(todo)
	helper.LogAndPanicError(validationError)

	todoModel := models.Todo{
		Description: todo.Description,
		Priority:    todo.Priority,
	}

	t.TodoRepository.Update(todoModel)
}

func (t *TodoServiceImpl) Delete(todoId int) {
	t.TodoRepository.Delete(todoId)
}
