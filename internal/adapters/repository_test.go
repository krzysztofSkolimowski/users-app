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
			repository := NewRepository(integrationTestsRepoConfig)
			repository.flush()
			for _, user := range tt.existingUsers {
				repository.insert(user)
			}

			err := repository.AddUser(tt.user)
			assert.Equal(t, tt.expectedErr, err)

			usersInRepo, _ := repository.allUsers()
			assert.Equal(t, tt.expectedUsers, usersInRepo)

		})
	}
}
