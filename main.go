package main

import (
	"net/http"
	"todo-crud/helper"
	"todo-crud/routes"

	"github.com/rs/zerolog/log"
)

func main() {
	routes := routes.InitRoutes()

	server := http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	log.Info().Msg("Server has started!")
	err := server.ListenAndServe()

	helper.LogAndPanicError(err)
}
