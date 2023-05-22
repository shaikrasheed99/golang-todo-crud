package main

import (
	"net/http"
	"todo-crud/config"
	"todo-crud/controllers"
	"todo-crud/helper"
	"todo-crud/models"
	"todo-crud/repository"
	"todo-crud/routes"
	"todo-crud/service"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("todos").AutoMigrate(&models.Todo{})

	todoRepository := repository.NewTodoRepositoryImpl(db)

	todoService := service.NewTodoServiceImpl(todoRepository, validate)

	todoController := controllers.NewTodoController(todoService)

	routes := routes.InitRoutes(todoController)

	server := http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	log.Info().Msg("Server has started!")
	err := server.ListenAndServe()

	helper.LogAndPanicError(err)
}
