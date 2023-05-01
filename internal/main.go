package main

import (
	"net/http"
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

	// todo - add a proper repository
	//mockRepo := adapters.NewMockRepo()

	repo := adapters.NewRepository()

	// todo - add a proper dependency injection framework google wire or something
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
