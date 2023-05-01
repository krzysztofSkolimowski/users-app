package adapters

import (
	"time"
	"users-app/domain"

	"github.com/google/uuid"
)

type UserDTO struct {
	ID           uuid.UUID `db:"id"`
	FirstName    string    `db:"first_name"`
	LastName     string    `db:"last_name"`
	Nickname     string    `db:"nickname"`
	PasswordHash string    `db:"password_hash"`
	Email        string    `db:"email"`
	Country      string    `db:"country"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func toDomainUsers(users []UserDTO) []domain.User {
	result := make([]domain.User, len(users))
	for i, user := range users {
		result[i] = toDomain(user)
	}

	return result
}

type MockRepo struct {
	users map[uuid.UUID]UserDTO
}

func (m *MockRepo) AddUser(user domain.User) (domain.UserID, error) {
	_, ok := m.users[user.ID]
	if ok {
		return domain.UserID{}, domain.ErrUserAlreadyExists
	}

	m.users[user.ID] = fromDomain(user)
	return user.ID, nil
}

func fromDomain(user domain.User) UserDTO {
	return UserDTO{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Nickname:     user.Nickname,
		PasswordHash: user.PasswordHash,
		Email:        user.Email,
		Country:      user.Country,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func (m *MockRepo) UpdateUser(user domain.User) error {
	_, ok := m.users[user.ID]
	if !ok {
		return domain.ErrUserNotFound
	}

	m.users[user.ID] = fromDomain(user)
	return nil
}

func (m *MockRepo) RemoveUser(id uuid.UUID) error {
	_, ok := m.users[id]
	if !ok {
		return domain.ErrUserNotFound
	}

	delete(m.users, id)
	return nil
}

func (m *MockRepo) Users(filter string) ([]domain.User, error) {
	var users []domain.User
	for _, user := range m.users {
		users = append(users, toDomain(user))
	}

	return users, nil
}

func toDomain(user UserDTO) domain.User {
	return domain.User{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Nickname:     user.Nickname,
		PasswordHash: user.PasswordHash,
		Email:        user.Email,
		Country:      user.Country,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func NewMockRepo() *MockRepo {
	return &MockRepo{
		users: make(map[domain.UserID]UserDTO),
	}
}
