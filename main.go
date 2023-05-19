package main

import (
	"todo-crud/routes"
)

func main() {
	r := routes.InitRoutes()

	r.Run()
}
