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
