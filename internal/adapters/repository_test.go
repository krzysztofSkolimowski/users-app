package adapters

import (
	"testing"
	"users-app/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var integrationTestsRepoConfig = RepoConfig{
	Host:     "localhost",
	Database: "users",
	User:     "postgres",
	Password: "password",
}

func setupRepo(existingUsers []domain.User) repository {
	repo := NewRepository(integrationTestsRepoConfig)
	repo.flush()
	for _, user := range existingUsers {
		repo.insert(user)
	}
	return repo
}

func Test_repository_AddUser(t *testing.T) {
	tests := []struct {
		name string
		user domain.User

		existingUsers []domain.User

		expectedErr   error
		expectedUsers []domain.User
	}{
		{
			name:          "user_gets_added_to_empty_repository",
			user:          domain.User{ID: uuid.MustParse("5f5d5ef5-5eb5-5cb5-b5d5-5f5d5ef5eb5c")},
			expectedErr:   nil,
			expectedUsers: []domain.User{{ID: uuid.MustParse("5f5d5ef5-5eb5-5cb5-b5d5-5f5d5ef5eb5c")}},
		},
		{
			name:          "user_does_not_get_added_if_already_exists",
			existingUsers: []domain.User{{ID: uuid.MustParse("5f5d5ef5-5eb5-5cb5-b5d5-5f5d5ef5eb5c")}},

			user:          domain.User{ID: uuid.MustParse("5f5d5ef5-5eb5-5cb5-b5d5-5f5d5ef5eb5c")},
			expectedErr:   domain.ErrUserAlreadyExists,
			expectedUsers: []domain.User{{ID: uuid.MustParse("5f5d5ef5-5eb5-5cb5-b5d5-5f5d5ef5eb5c")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := setupRepo(tt.existingUsers)

			err := repo.AddUser(tt.user)
			assert.Equal(t, tt.expectedErr, err)

			usersInRepo, _ := repo.allUsers()
			assert.Equal(t, tt.expectedUsers, usersInRepo)
		})
	}
}

func Test_repository_Users(t *testing.T) {
	stringPtr := func(s string) *string { return &s }
	// fixtures
	uuid_1 := uuid.MustParse("5f5d5ef5-5eb5-5cb5-b5d5-5f5d5ef5eb5c")
	uuid_2 := uuid.MustParse("7a13e2ff-2c47-4f16-9c35-8e24abddc0ea")
	uuid_3 := uuid.MustParse("c95b7c8a-9e64-4e22-81df-a2e8fcb30c81")

	user_1 := domain.User{
		ID: uuid_1, FirstName: "John", LastName: "Doe",
		Nickname: "johndoe", Email: "john.doe@email.com", Country: "US",
	}
	user_2 := domain.User{
		ID: uuid_2, FirstName: "Jane", LastName: "Doe",
		Nickname: "janedoe", Email: "jane.doe@emai.com", Country: "US",
	}
	user_3 := domain.User{
		ID: uuid_3, FirstName: "John", LastName: "Stones",
		Nickname: "jackstones", Email: "jack.stones@email.com", Country: "UK",
	}

	tests := []struct {
		name       string
		filter     domain.Filter
		pagination domain.Pagination

		existingUsers []domain.User

		expected    []domain.User
		expectedErr error
	}{
		{
			name:          "users_get_returned",
			existingUsers: []domain.User{user_1, user_2},
			expected:      []domain.User{user_1, user_2},
		},
		{
			name:          "results_are_paginated_if_pagination_is_set",
			existingUsers: []domain.User{user_1, user_2, user_3},
			pagination:    domain.NewPagination(2, 0),
			expected:      []domain.User{user_1, user_2},
		},
		{
			name:          "first_name_filter_returns_existing_user",
			existingUsers: []domain.User{user_1, user_2},
			filter:        domain.Filter{FirstName: stringPtr("John")},
			expected:      []domain.User{user_1},
		},
		{
			name: "last_name_filter_returns_existing_users",
			existingUsers: []domain.User{
				user_1, user_2, user_3,
			},
			filter:   domain.Filter{LastName: stringPtr("Doe")},
			expected: []domain.User{user_1, user_2},
		},
		{
			name: "nickname_filter_returns_existing_user",
			existingUsers: []domain.User{
				user_1, user_2,
			},
			filter:   domain.Filter{Nickname: stringPtr("johndoe")},
			expected: []domain.User{user_1},
		},
		{
			name: "email_filter_returns_existing_user",
			existingUsers: []domain.User{
				user_1, user_2,
			},
			filter:   domain.Filter{Email: stringPtr("john.doe@email.com")},
			expected: []domain.User{user_1},
		},
		{
			name: "country_filter_returns_existing_user",
			existingUsers: []domain.User{
				user_1, user_2,
			},
			filter:   domain.Filter{Country: stringPtr("US")},
			expected: []domain.User{user_1, user_2},
		},
		{
			name: "multiple_filters_return_correct_results",
			existingUsers: []domain.User{
				user_1, user_2, user_3,
			},
			filter: domain.Filter{
				FirstName: stringPtr("John"),
				Country:   stringPtr("UK"),
			},
			expected: []domain.User{user_3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			repo := setupRepo(tt.existingUsers)

			ret, err := repo.Users(tt.filter, tt.pagination)

			assert.Equal(t, tt.expectedErr, err)
			assert.EqualValues(t, tt.expected, ret)
		})
	}
}
