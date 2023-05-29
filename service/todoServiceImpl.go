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

func (ts *TodoServiceImpl) GetAll() []response.TodoResponse {
	todos := ts.TodoRepository.FindAll()

	var todosResponse []response.TodoResponse

	for _, value := range todos {
		todo := response.TodoResponse(value)
		todosResponse = append(todosResponse, todo)
	}

	return todosResponse
}

func isTodoEmpty(todo models.Todo) bool {
	return (todo.Id == 0 && todo.Description == "" && todo.Priority == "")
}

func (ts *TodoServiceImpl) GetById(todoId int) response.TodoResponse {
	todo := ts.TodoRepository.FindById(todoId)

	if isTodoEmpty(todo) {
		helper.LogAndPanicError(errors.New("there is no todo with this id"))
	}

	todoResponse := response.TodoResponse(todo)

	return todoResponse
}

func (ts *TodoServiceImpl) Create(requestTodo request.CreateTodoRequest) response.TodoResponse {
	validationError := ts.Validate.Struct(requestTodo)
	helper.LogAndPanicError(validationError)

	newTodo := models.Todo{
		Description: requestTodo.Description,
		Priority:    requestTodo.Priority,
	}

	createdTodo := ts.TodoRepository.Save(newTodo)

	todoResponse := response.TodoResponse(createdTodo)

	return todoResponse
}

func (ts *TodoServiceImpl) Update(todoId int, requestTodo request.UpdateTodoRequest) response.TodoResponse {
	validationError := ts.Validate.Struct(requestTodo)
	helper.LogAndPanicError(validationError)

	todoById := ts.TodoRepository.FindById(todoId)

	if isTodoEmpty(todoById) {
		helper.LogAndPanicError(errors.New("there is no todo with this id"))
	}

	newTodo := models.Todo{
		Id:          todoById.Id,
		Description: requestTodo.Description,
		Priority:    requestTodo.Priority,
	}

	ts.TodoRepository.Update(todoId, newTodo)

	todoResponse := response.TodoResponse(newTodo)

	return todoResponse
}

func (ts *TodoServiceImpl) Delete(todoId int) {
	todoById := ts.TodoRepository.FindById(todoId)

	if isTodoEmpty(todoById) {
		helper.LogAndPanicError(errors.New("there is no todo with this id"))
	}

	ts.TodoRepository.Delete(todoId)
}
