package grpc

import (
	"context"
	"users-app/domain"
	"users-app/gen/grpc"
	"users-app/service"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UsersServer struct {
	users_app.UnimplementedUsersServer
	queryService   service.UsersQueryService
	commandService service.UsersCommandService
}

func NewGRPCServer(queryService service.UsersQueryService, commandService service.UsersCommandService) *UsersServer {
	return &UsersServer{
		queryService:   queryService,
		commandService: commandService,
	}
}

func (s *UsersServer) HealthCheck(ctx context.Context, in *empty.Empty) (*users_app.HealthCheckResponse, error) {
	return &users_app.HealthCheckResponse{Status: "OK"}, nil
}

func (s *UsersServer) GetUsers(ctx context.Context, in *users_app.GetUsersRequest) (*users_app.GetUsersResponse, error) {
	users, err := s.queryService.Users(ctx, parseFilter(in.GetFilter()), parsePagination(in.GetPagination()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error")
	}

	return getUsersResponse(users), nil
}

func (s *UsersServer) CreateUser(ctx context.Context, in *users_app.CreateUserRequest) (*users_app.User, error) {
	user, err := s.commandService.AddUser(ctx, service.AddUserCommand{
		FirstName: in.GetFirstName(),
		LastName:  in.GetLastName(),
		Nickname:  in.GetNickname(),
		Password:  in.GetPassword(),
		Email:     in.GetEmail(),
		Country:   in.GetCountry(),
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error")
	}

	return toGRPCUserResponse(user), nil
}

func (s *UsersServer) ModifyUser(ctx context.Context, in *users_app.ModifyUserRequest) (*users_app.ModifyUserResponse, error) {
	id, err := domain.ParseID(in.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id")
	}

	err = s.commandService.ModifyUser(ctx, modifyCommand(id, in))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error")
	}

	return &users_app.ModifyUserResponse{Status: "OK"}, nil
}

func (s *UsersServer) DeleteUser(ctx context.Context, in *users_app.DeleteUserRequest) (*empty.Empty, error) {
	id, err := domain.ParseID(in.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id")
	}

	err = s.commandService.DeleteUser(ctx, service.DeleteUserCommand{id})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error")
	}

	return &empty.Empty{}, nil
}
