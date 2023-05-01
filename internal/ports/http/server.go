package http

import (
	"net/http"
	"users-app/domain"
	"users-app/gen/api"
	"users-app/service"

	"github.com/go-chi/render"
	"github.com/labstack/gommon/log"
)

type Server struct {
	queryService   service.UsersQueryService
	commandService service.UsersCommandService
}

func (h Server) GetUsers(w http.ResponseWriter, r *http.Request, params api.GetUsersParams) {
	users, err := h.queryService.Users(r.Context(), filterFromParams(params), paginationFromParams(params))
	if err != nil {
		log.Error(err)
		render.Respond(w, r, api.Error{Code: http.StatusInternalServerError, Message: "internal server error"})
		return
	}

	render.Respond(w, r, api.Users{
		Users: usersListFromDomain(users),
	})

	return

}

func NewHttpServer(queries service.UsersQueryService, commands service.UsersCommandService) Server {

	return Server{queries, commands}
}

func (h Server) GetHealth(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, "ok")
	return
}

func (h Server) PostUsers(w http.ResponseWriter, r *http.Request) {
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

func (h Server) DeleteUsersUserID(w http.ResponseWriter, r *http.Request, userID string) {
	id, err := domain.ParseID(userID)
	if err != nil {
		log.Error(err)
		render.Respond(w, r, api.Error{Code: http.StatusBadRequest, Message: "invalid user id"})
		return
	}

	h.commandService.DeleteUser(r.Context(), service.DeleteUserCommand{ID: id})
	render.Respond(w, r, "ok")
	return
}

func (h Server) PatchUsersUserID(w http.ResponseWriter, r *http.Request, userID string) {
	render.Respond(w, r, "ok")
	return
}
