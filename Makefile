
generate:
	oapi-codegen -generate types  -package api api/users.yml > internal/gen/api/http_api_types.go
	oapi-codegen -generate chi-server -package api api/users.yml > internal/gen/api/http_server.go

up:
	docker compose up