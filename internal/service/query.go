package service

import (
	"context"
	"users-app/domain"
)

type UsersQueryService interface {
	Users(ctx context.Context, filter string) ([]domain.User, error)
}

type UserQueryService struct {
	userRepository domain.Repository
}

func NewUserQueryService(repo domain.Repository) *UserQueryService {
	return &UserQueryService{repo}
}

func (u UserQueryService) Users(ctx context.Context, filter string) ([]domain.User, error) {
	return u.userRepository.Users(filter)

}
