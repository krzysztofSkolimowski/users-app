package ports

import (
	"net/http"
	"users-app/gen/api"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/go-chi/render"
	"github.com/go-openapi/swag"
	"github.com/google/uuid"
)

type HttpServer struct {
	// todo - remove json raw message and move to shared package
	swaggerJson []byte
}

func NewHttpServer(swaggerDocPath string) HttpServer {
	// todo - move somewhere else
	// read swagger dock
	spec, err := swag.YAMLDoc(swaggerDocPath)
	if err != nil {
		panic(err)
	}

	return HttpServer{
		swaggerJson: spec,
	}
}

func (h HttpServer) GetHealth(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, "ok")
	return
}

func (h HttpServer) GetUsers(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, api.Users{
		Users: []api.User{
			{Country: "United States", Email: "john.doe@example.com", FirstName: "John", Id: openapi_types.UUID{}, LastName: "Doe", Nickname: "JD", UpdatedAt: nil},
			{Country: "France", Email: "jane.smith@example.com", FirstName: "Jane", Id: openapi_types.UUID{}, LastName: "Smith", Nickname: "JS", UpdatedAt: nil},
			{Country: "Germany", Email: "michael.chang@example.com", FirstName: "Michael", Id: openapi_types.UUID{}, LastName: "Chang", Nickname: "MC", UpdatedAt: nil},
			{Country: "United Kingdom", Email: "susan.jones@example.com", FirstName: "Susan", Id: openapi_types.UUID{}, LastName: "Jones", Nickname: "SJ", UpdatedAt: nil},
			{Country: "Canada", Email: "james.harris@example.com", FirstName: "James", Id: openapi_types.UUID{}, LastName: "Harris", Nickname: "JH", UpdatedAt: nil},
		},
	})
	return
}

func (h HttpServer) PostUsers(w http.ResponseWriter, r *http.Request) {
	id := uuid.New()
	user := api.User{Country: "United Kingdom", Email: "susan.jones@example.com", FirstName: "Susan", Id: id, LastName: "Jones", Nickname: "SJ", UpdatedAt: nil}
	render.Respond(w, r, user)
	return
}

func (h HttpServer) DeleteUsersUserID(w http.ResponseWriter, r *http.Request, userID string) {
	render.Respond(w, r, "ok")
	return
}

func (h HttpServer) PatchUsersUserID(w http.ResponseWriter, r *http.Request, userID string) {
	render.Respond(w, r, "ok")
	return
}
