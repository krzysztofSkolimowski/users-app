package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"users-app/adapters"
	"users-app/gen/api"
	users_app "users-app/gen/grpc"
	ports_grpc "users-app/ports/grpc"
	ports "users-app/ports/http"
	"users-app/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	repo := adapters.NewRepository(
		adapters.RepoConfig{
			Host:     os.Getenv("DB_HOST"),
			Database: os.Getenv("POSTGRES_DB"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
		},
	)
	querySvc := service.NewUserQueryService(repo)

	redis := adapters.NewPubSub(adapters.RedisConfig{
		Host:          getEnvString("REDIS_HOST", "redis"),
		Port:          getEnvString("REDIS_PORT", "6379"),
		Password:      getEnvString("REDIS_PASSWORD", ""), // todo - no password set for now
		DB:            getEnvInt("REDIS_DB", 0),
		EventsChannel: getEnvString("REDIS_EVENTS_CHANNEL", "events"),
	})

	commandSvcBase := service.NewUserCommandService(repo)
	eventsLogFilePath := getEnvString("EVENTS_LOG_FILE_PATH", "../logs/events.log")

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	commandSvcLogging := service.NewCommandLoggingWrapper(logger, commandSvcBase)

	commandSvc := service.NewCommandEventsWrapper(redis, commandSvcLogging, adapters.NewEventLogger(eventsLogFilePath))

	if getEnvBool("RUN_HTTP", true) {
		go runHTTPServer(querySvc, commandSvc)
	}

	if getEnvBool("RUN_GRPC", true) {
		go runGRPCServer(querySvc, commandSvc)
	}

	select {}
}

func runGRPCServer(querySvc service.UsersQueryService, commandSvc service.UsersCommandService) {
	grpcServer := grpc.NewServer()

	usersServer := ports_grpc.NewGRPCServer(querySvc, commandSvc)
	users_app.RegisterUsersServer(grpcServer, usersServer)

	port := getEnvString("PORT_GRPC", "50051")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen on port: %v, error: %v", port, err)
	}

	log.Printf("gRPC server listening on :%s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func runHTTPServer(querySvc service.UsersQueryService, commandSvc service.UsersCommandService) {
	router := chi.NewRouter()

	httpServer := ports.NewHttpServer(querySvc, commandSvc)
	handler := api.HandlerWithOptions(httpServer, api.ChiServerOptions{
		BaseRouter: router,
		Middlewares: []api.MiddlewareFunc{
			middleware.RequestID,
			httplog.RequestLogger(httplog.NewLogger("users-app", httplog.Options{
				LogLevel:        getEnvString("LOG_LEVEL", "info"),
				LevelFieldName:  "level",
				JSON:            getEnvBool("LOG_JSON", true),
				TimeFieldFormat: "2006-01-02T15:04:05.000Z07:00",
				TimeFieldName:   "timestamp",
			})),
			middleware.Recoverer,
		},
	})

	port := getEnvString("PORT_HTTP", "8080")
	log.Printf("HTTP server listening on :%s", port)
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatalf("failed to listen and serve: %v", err)
	}
}

func getEnvString(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return value
}

func getEnvBool(key string, defaultValue bool) bool {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return value == "true"
}

func getEnvInt(s string, i int) int {
	value, ok := os.LookupEnv(s)
	if !ok {
		return i
	}

	v, err := strconv.Atoi(strings.TrimSpace(value))
	if err != nil {
		panic(err)
	}

	return v
}
