package main

import (
	"net/http"
	"users-app/gen/api"
	"users-app/ports"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	runServer()
}

func runServer() {

	router := chi.NewRouter()

	// todo - export to env

	httpServer := ports.NewHttpServer("../api/users.yml")

	handler := api.HandlerWithOptions(httpServer, api.ChiServerOptions{
		BaseRouter: router,
		Middlewares: []api.MiddlewareFunc{
			middleware.Logger,
		},
	})

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}
}
