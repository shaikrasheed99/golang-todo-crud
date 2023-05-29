package service

import (
	"errors"
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

func isTodoEmpty(todo models.Todo) bool {
	return (todo.Id == 0 && todo.Description == "" && todo.Priority == "")
}

func (t *TodoServiceImpl) GetById(todoId int) response.TodoResponse {
	todo := t.TodoRepository.FindById(todoId)

	if isTodoEmpty(todo) {
		helper.LogAndPanicError(errors.New("there is no todo with this id"))
	}

	todoResponse := response.TodoResponse{
		Id:          todo.Id,
		Description: todo.Description,
		Priority:    todo.Priority,
	}

	return todoResponse
}

func (t *TodoServiceImpl) Create(todoRequest request.CreateTodoRequest) response.TodoResponse {
	validationError := t.Validate.Struct(todoRequest)
	helper.LogAndPanicError(validationError)

	newTodo := models.Todo{
		Description: todoRequest.Description,
		Priority:    todoRequest.Priority,
	}

	createdTodo := t.TodoRepository.Save(newTodo)

	todoResponse := response.TodoResponse{
		Id:          createdTodo.Id,
		Description: createdTodo.Description,
		Priority:    createdTodo.Priority,
	}

	return todoResponse
}

func (t *TodoServiceImpl) Update(todoId int, requestTodo request.UpdateTodoRequest) {
	validationError := t.Validate.Struct(requestTodo)
	helper.LogAndPanicError(validationError)

	todoById := t.TodoRepository.FindById(todoId)

	if isTodoEmpty(todoById) {
		helper.LogAndPanicError(errors.New("there is no todo with this id"))
	}

	newTodo := models.Todo{
		Id:          todoById.Id,
		Description: requestTodo.Description,
		Priority:    requestTodo.Priority,
	}

	t.TodoRepository.Update(todoId, newTodo)
}

func (t *TodoServiceImpl) Delete(todoId int) {
	todoById := t.TodoRepository.FindById(todoId)

	if isTodoEmpty(todoById) {
		helper.LogAndPanicError(errors.New("there is no todo with this id"))
	}

	t.TodoRepository.Delete(todoId)
}
