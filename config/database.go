package config

import (
	"fmt"
	"todo-crud/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "todos"
)

func InitDatabase() *gorm.DB {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	helper.LogAndPanicError(err)

	return db
}
