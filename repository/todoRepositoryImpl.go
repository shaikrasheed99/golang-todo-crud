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

func (t *TodoRepositoryImpl) FindAll() []models.Todo {
	var todos []models.Todo
	result := t.Db.Find(&todos)
	helper.LogAndPanicError(result.Error)
	return todos
}

func (t *TodoRepositoryImpl) Save(todo models.Todo) {
	result := t.Db.Create(&todo)
	helper.LogAndPanicError(result.Error)
}

func (t *TodoRepositoryImpl) Update(todo models.Todo) {
	result := t.Db.Model(&todo).Updates(todo)
	helper.LogAndPanicError(result.Error)
}

func (t *TodoRepositoryImpl) Delete(todoId int) {
	var todo models.Todo
	result := t.Db.Where("id = ?", todoId).Delete(&todo)
	helper.LogAndPanicError(result.Error)
}
