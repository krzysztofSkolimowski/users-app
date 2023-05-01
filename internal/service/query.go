package service

import (
	"context"
	"users-app/domain"
)

type UsersQueryService interface {
	Users(context.Context, domain.Filter, domain.Pagination) ([]domain.User, error)
}

type UserQueryService struct {
	userRepository domain.Repository
}

func NewUserQueryService(repo domain.Repository) *UserQueryService {
	return &UserQueryService{repo}
}

func (u UserQueryService) Users(ctx context.Context, f domain.Filter, p domain.Pagination) ([]domain.User, error) {
	return u.userRepository.Users(f, p)
}
