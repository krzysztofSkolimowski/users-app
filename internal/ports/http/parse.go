package http

import (
	"users-app/domain"
	"users-app/gen/api"

	"github.com/oapi-codegen/runtime/types"
)

func filterFromParams(params api.GetUsersParams) domain.Filter {
	mail := ""
	if params.Email != nil {
		mail = string(*params.Email)
	}

	return domain.NewFilter(
		nilSafeString(params.FirstName),
		nilSafeString(params.LastName),
		nilSafeString(params.Nickname),
		mail,
		nilSafeString(params.Country),
	)

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

func nilSafeString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
