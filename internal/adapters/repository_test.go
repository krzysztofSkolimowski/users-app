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
	uuid1 := uuid.MustParse("5f5d5ef5-5eb5-5cb5-b5d5-5f5d5ef5eb5c")
	uuid2 := uuid.MustParse("7a13e2ff-2c47-4f16-9c35-8e24abddc0ea")
	uuid3 := uuid.MustParse("c95b7c8a-9e64-4e22-81df-a2e8fcb30c81")

	user1 := domain.User{
		ID: uuid1, FirstName: "John", LastName: "Doe",
		Nickname: "johndoe", Email: "john.doe@email.com", Country: "US",
	}
	user2 := domain.User{
		ID: uuid2, FirstName: "Jane", LastName: "Doe",
		Nickname: "janedoe", Email: "jane.doe@emai.com", Country: "US",
	}
	user3 := domain.User{
		ID: uuid3, FirstName: "John", LastName: "Stones",
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
			existingUsers: []domain.User{user1, user2},
			expected:      []domain.User{user1, user2},
		},
		{
			name:          "results_are_paginated_if_pagination_is_set",
			existingUsers: []domain.User{user1, user2, user3},
			pagination:    domain.NewPagination(2, 0),
			expected:      []domain.User{user1, user2},
		},
		{
			name:          "first_name_filter_returns_existing_user",
			existingUsers: []domain.User{user1, user2},
			filter:        domain.Filter{FirstName: stringPtr("John")},
			expected:      []domain.User{user1},
		},
		{
			name: "last_name_filter_returns_existing_users",
			existingUsers: []domain.User{
				user1, user2, user3,
			},
			filter:   domain.Filter{LastName: stringPtr("Doe")},
			expected: []domain.User{user1, user2},
		},
		{
			name: "nickname_filter_returns_existing_user",
			existingUsers: []domain.User{
				user1, user2,
			},
			filter:   domain.Filter{Nickname: stringPtr("johndoe")},
			expected: []domain.User{user1},
		},
		{
			name: "email_filter_returns_existing_user",
			existingUsers: []domain.User{
				user1, user2,
			},
			filter:   domain.Filter{Email: stringPtr("john.doe@email.com")},
			expected: []domain.User{user1},
		},
		{
			name: "country_filter_returns_existing_user",
			existingUsers: []domain.User{
				user1, user2,
			},
			filter:   domain.Filter{Country: stringPtr("US")},
			expected: []domain.User{user1, user2},
		},
		{
			name: "multiple_filters_return_correct_results",
			existingUsers: []domain.User{
				user1, user2, user3,
			},
			filter: domain.Filter{
				FirstName: stringPtr("John"),
				Country:   stringPtr("UK"),
			},
			expected: []domain.User{user3},
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

func Test_repository_RemoveUser(t *testing.T) {
	// fixtures
	uuid1 := uuid.MustParse("5f5d5ef5-5eb5-5cb5-b5d5-5f5d5ef5eb5c")
	uuid2 := uuid.MustParse("7a13e2ff-2c47-4f16-9c35-8e24abddc0ea")
	uuid3 := uuid.MustParse("c95b7c8a-9e64-4e22-81df-a2e8fcb30c81")

	user1 := domain.User{
		ID: uuid1, FirstName: "John", LastName: "Doe", Nickname: "johndoe", Email: "john@doe.com", Country: "US",
	}
	user2 := domain.User{
		ID: uuid2, FirstName: "Jane", LastName: "Doe", Nickname: "janedoe", Email: "jane@doe.com", Country: "US",
	}
	user3 := domain.User{
		ID: uuid3, FirstName: "John", LastName: "Stones", Nickname: "jstones", Email: "jack@stones.com", Country: "UK",
	}

	tests := []struct {
		name string
		id   domain.UserID

		existingUsers []domain.User

		expectedErr   error
		expectedUsers []domain.User
	}{
		{name: "one_in_repo", id: uuid1, existingUsers: []domain.User{user1}, expectedUsers: []domain.User{}},
		{name: "many_in_repo", id: uuid1, existingUsers: []domain.User{user1, user2, user3}, expectedUsers: []domain.User{user2, user3}},
		{name: "user_does_not_exist", id: uuid1, expectedUsers: []domain.User{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := setupRepo(tt.existingUsers)

			err := repo.RemoveUser(tt.id)
			assert.Equal(t, tt.expectedErr, err)

			usersInRepo, _ := repo.allUsers()
			assert.Equal(t, tt.expectedUsers, usersInRepo)
		})
	}
}

func Test_repository_ModifyUser(t *testing.T) {
	uuid1 := uuid.MustParse("5f5d5ef5-5eb5-5cb5-b5d5-5f5d5ef5eb5c")
	uuid2 := uuid.MustParse("7a13e2ff-2c47-4f16-9c35-8e24abddc0ea")
	user1 := domain.User{
		ID: uuid1, FirstName: "John", LastName: "Doe", Nickname: "johndoe", Email: "john@doe.com", Country: "US",
	}

	user1NameModified := user1
	user1NameModified.FirstName = "Alex"

	user1SurnameModified := user1
	user1SurnameModified.LastName = "Smith"

	user1NicknameModified := user1
	user1NicknameModified.Nickname = "alex"

	user1EmailModified := user1
	user1EmailModified.Email = "new@email.com"

	user1CountryModified := user1
	user1CountryModified.Country = "UK"

	user2 := domain.User{
		ID: uuid2, FirstName: "Jane", LastName: "Doe", Nickname: "janedoe", Email: "jane@doe.com", Country: "US",
	}

	tests := []struct {
		name   string
		id     domain.UserID
		fields domain.Fields

		existingUsers []domain.User

		expectedErr   error
		expectedUsers []domain.User
	}{
		{
			name:          "user_does_not_exist",
			id:            uuid1,
			fields:        domain.Fields{"first_name": "Alex"},
			expectedErr:   domain.ErrUserNotFound,
			expectedUsers: []domain.User{},
		},
		{
			name: "user_exists_first_name_gets_modified",
			id:   uuid1,
			fields: domain.Fields{
				"first_name": "Alex",
			},
			existingUsers: []domain.User{user1},
			expectedUsers: []domain.User{user1NameModified},
		},
		{
			name: "user_exists_last_name_gets_modified",
			id:   uuid1,
			fields: domain.Fields{
				"last_name": "Smith",
			},
			existingUsers: []domain.User{user1, user2},
			expectedUsers: []domain.User{
				user1SurnameModified,
				user2,
			},
		},
		{
			name: "user_exists_nickname_gets_modified",
			id:   uuid1,
			fields: domain.Fields{
				"nickname": "alex",
			},
			existingUsers: []domain.User{user1, user2},
			expectedUsers: []domain.User{
				user1NicknameModified,
				user2,
			},
		},
		{
			name: "user_exists_email_gets_modified",
			id:   uuid1,
			fields: domain.Fields{
				"email": "new@email.com",
			},
			existingUsers: []domain.User{user1, user2},
			expectedUsers: []domain.User{
				user1EmailModified,
				user2,
			},
		},
		{
			name: "user_exists_country_gets_modified",
			id:   uuid1,
			fields: domain.Fields{
				"country": "UK",
			},
			existingUsers: []domain.User{user1, user2},
			expectedUsers: []domain.User{
				user1CountryModified,
				user2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := setupRepo(tt.existingUsers)

			err := repo.ModifyUser(tt.id, tt.fields)
			assert.Equal(t, tt.expectedErr, err)

			usersInRepo, _ := repo.allUsers()
			assert.ElementsMatch(t, tt.expectedUsers, usersInRepo)
		})
	}
}
