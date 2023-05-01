package ports

import (
	"net/http"
	"users-app/domain"
	"users-app/gen/api"
	"users-app/service"

	types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/go-chi/render"
	"github.com/labstack/gommon/log"
)

type HttpServer struct {
	queryService   service.UsersQueryService
	commandService service.UsersCommandService
}

func NewHttpServer(queries service.UsersQueryService, commands service.UsersCommandService) HttpServer {

	return HttpServer{queries, commands}
}

func (h HttpServer) GetHealth(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, "ok")
	return
}

func (h HttpServer) GetUsers(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, api.Users{
		Users: []api.User{
			{Country: "United States", Email: "john.doe@example.com", FirstName: "John", Id: types.UUID{}, LastName: "Doe", Nickname: "JD", UpdatedAt: nil},
			{Country: "France", Email: "jane.smith@example.com", FirstName: "Jane", Id: types.UUID{}, LastName: "Smith", Nickname: "JS", UpdatedAt: nil},
			{Country: "Germany", Email: "michael.chang@example.com", FirstName: "Michael", Id: types.UUID{}, LastName: "Chang", Nickname: "MC", UpdatedAt: nil},
			{Country: "United Kingdom", Email: "susan.jones@example.com", FirstName: "Susan", Id: types.UUID{}, LastName: "Jones", Nickname: "SJ", UpdatedAt: nil},
			{Country: "Canada", Email: "james.harris@example.com", FirstName: "James", Id: types.UUID{}, LastName: "Harris", Nickname: "JH", UpdatedAt: nil},
		},
	})
	return
}

func (h HttpServer) PostUsers(w http.ResponseWriter, r *http.Request) {
	postUser := api.PostUser{}
	err := render.Decode(r, &postUser)
	if err != nil {
		// todo - implement proper logging
		log.Error(err)
		render.Respond(w, r, api.Error{Code: http.StatusBadRequest, Message: "invalid request"})
		return
	}

	user, err := h.commandService.AddUser(r.Context(), service.AddUserCommand{
		FirstName: postUser.FirstName,
		LastName:  postUser.LastName,
		Nickname:  postUser.Nickname,
		Password:  postUser.Password,
		Email:     string(postUser.Email),
		Country:   postUser.Country,
	})
	if err != nil {
		log.Error(err)
		// todo - implement proper error handling
		render.Respond(w, r, api.Error{Code: http.StatusInternalServerError, Message: "internal server error"})
		return
	}

	render.Respond(w, r, toUserResponse(user))
	return
}

func toUserResponse(user domain.User) api.User {
	return api.User{
		Country:   user.Country,
		Email:     types.Email(user.Email),
		FirstName: user.FirstName,
		Id:        user.ID,
		LastName:  user.LastName,
		Nickname:  user.Nickname,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
	}
}

func (h HttpServer) DeleteUsersUserID(w http.ResponseWriter, r *http.Request, userID string) {
	render.Respond(w, r, "ok")
	return
}

func (h HttpServer) PatchUsersUserID(w http.ResponseWriter, r *http.Request, userID string) {
	render.Respond(w, r, "ok")
	return
}
