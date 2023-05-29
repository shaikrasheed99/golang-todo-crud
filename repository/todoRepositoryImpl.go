package repository

import (
	"todo-crud/helper"
	"todo-crud/models"

	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	Db *gorm.DB
}

func NewTodoRepositoryImpl(Db *gorm.DB) TodoRepository {
	return &TodoRepositoryImpl{Db: Db}
}

func (tr *TodoRepositoryImpl) FindAll() []models.Todo {
	var todos []models.Todo
	result := tr.Db.Find(&todos)
	helper.LogAndPanicError(result.Error)
	return todos
}

func (tr *TodoRepositoryImpl) FindById(todoId int) models.Todo {
	var todo models.Todo
	result := tr.Db.Where("id = ?", todoId).Find(&todo)
	helper.LogAndPanicError(result.Error)
	return todo
}

func (tr *TodoRepositoryImpl) Save(todo models.Todo) models.Todo {
	result := tr.Db.Create(&todo)
	helper.LogAndPanicError(result.Error)
	return todo
}

func (tr *TodoRepositoryImpl) Update(todoId int, newTodo models.Todo) {
	result := tr.Db.Where("id = ?", todoId).Updates(newTodo)
	helper.LogAndPanicError(result.Error)
}

func (tr *TodoRepositoryImpl) Delete(todoId int) {
	var todo models.Todo
	result := tr.Db.Where("id = ?", todoId).Delete(&todo)
	helper.LogAndPanicError(result.Error)
}
