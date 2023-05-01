
generate:
	oapi-codegen -generate types  -package api api/users.yml > internal/gen/api/http_api_types.go
	oapi-codegen -generate chi-server -package api api/users.yml > internal/gen/api/http_server.go
	protoc -I api --go_out=internal/gen/grpc --go_opt=paths=source_relative --go-grpc_out=internal/gen/grpc --go-grpc_opt=paths=source_relative api/users.proto

up:
	docker compose up

test:
	gotestsum -f testname ./internal/...

evans:
	evans  api/users.proto
