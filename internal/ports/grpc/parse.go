package grpc

import (
	"users-app/domain"
	users_app "users-app/gen/grpc"
	"users-app/service"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func parseFilter(in *users_app.Filter) domain.Filter {
	return domain.NewFilter(in.GetFirstName(), in.GetLastName(), in.GetNickname(), in.GetEmail(), in.GetCountry())
}

func parsePagination(pagination *users_app.Pagination) domain.Pagination {
	domain.NewPagination(int(pagination.GetLimit()), int(pagination.GetOffset()))

	return domain.Pagination{}
}

func toGRPCUserResponse(user domain.User) *users_app.User {
	return &users_app.User{
		Id:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Nickname:  user.Nickname,
		Email:     user.Email,
		Country:   user.Country,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func modifyCommand(id domain.UserID, in *users_app.ModifyUserRequest) service.ModifyUserCommand {
	ret := service.ModifyUserCommand{}
	ret.ID = id
	if in.FirstName != "" {
		ret.FirstName = stringPTR(in.FirstName)
	}
	if in.LastName != "" {
		ret.LastName = stringPTR(in.LastName)
	}
	if in.Nickname != "" {
		ret.Nickname = stringPTR(in.Nickname)
	}
	if in.Email != "" {
		ret.Email = stringPTR(in.Email)
	}
	if in.Country != "" {
		ret.Country = stringPTR(in.Country)
	}

	return ret
}

func getUsersResponse(users []domain.User) *users_app.GetUsersResponse {

	ret := &users_app.GetUsersResponse{}
	ret.Users = make([]*users_app.User, len(users))

	for i, user := range users {
		ret.Users[i] = toGRPCUserResponse(user)
	}

	return ret
}

func stringPTR(s string) *string {
	return &s
}
