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

// Users returns a list of users, filtered and paginated
// In case of no pagination passed, it will return domain.DefaultPagination (10 users, offset 0)
// In case of no filter passed, it will return all users (paginated) from the repository
func (u UserQueryService) Users(ctx context.Context, f domain.Filter, p domain.Pagination) ([]domain.User, error) {
	// todo - error handling
	return u.userRepository.Users(f, p)
}
