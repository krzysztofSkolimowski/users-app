package main

import (
	"net/http"
	"os"
	"users-app/adapters"
	"users-app/gen/api"
	"users-app/ports"
	"users-app/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
)

func main() {
	runServer()
}

func runServer() {
	router := chi.NewRouter()

	repo := adapters.NewRepository(
		adapters.RepoConfig{
			Host:     os.Getenv("DB_HOST"),
			Database: os.Getenv("POSTGRES_DB"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
		})

	// todo - add a proper dependency injection framework google wire or uber.fx
	querySvc := service.NewUserQueryService(repo)
	commandSvc := service.NewUserCommandService(repo)
	httpServer := ports.NewHttpServer(querySvc, commandSvc)

	handler := api.HandlerWithOptions(httpServer, api.ChiServerOptions{
		BaseRouter: router,
		Middlewares: []api.MiddlewareFunc{
			middleware.RequestID,
			httplog.RequestLogger(httplog.NewLogger("users-app", httplog.Options{
				LogLevel:        "info",
				LevelFieldName:  "level",
				JSON:            true,
				TimeFieldFormat: "2006-01-02T15:04:05.000Z07:00",
				TimeFieldName:   "timestamp",
			})),
			middleware.Recoverer,
		},
	})

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}
}
