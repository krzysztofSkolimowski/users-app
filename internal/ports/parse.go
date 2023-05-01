package ports

import (
	"users-app/domain"
	"users-app/gen/api"

	"github.com/deepmap/oapi-codegen/pkg/types"
)

func filterFromParams(params api.GetUsersParams) domain.Filter {
	ret := domain.Filter{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Nickname:  params.Nickname,
		Country:   params.Country,
	}

	if params.Email != nil {
		mail := string(*params.Email)
		ret.Email = &mail
	}

	return ret
}

func usersListFromDomain(users []domain.User) []api.User {
	usersApi := make([]api.User, len(users))

	for i, user := range users {
		usersApi[i] = toUserResponse(user)
	}

	return usersApi
}

func toUserResponse(user domain.User) api.User {
	return api.User{
		Country:   user.Country,
		Email:     types.Email(user.Email),
		FirstName: user.FirstName,
		Id:        user.ID,
		LastName:  user.LastName,
		Nickname:  user.Nickname,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func paginationFromParams(params api.GetUsersParams) domain.Pagination {
	limit, offset := 0, 0
	if params.Limit != nil {
		limit = int(*params.Limit)
	}
	if params.Offset != nil {
		offset = int(*params.Offset)
	}

	return domain.NewPagination(limit, offset)
}
