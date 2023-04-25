package adapters

import (
	"time"
	"users-app/internal/domain"
)

type UserDTO struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Nickname  string    `json:"nickname"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepo struct{}

func NewRepo() UserRepo {
	// init connection

	return UserRepo{}
}

func (u UserRepo) AddUser(user domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) UpdateUser(user domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) RemoveUser(id string) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) Users(filter string) ([]domain.User, error) {
	//TODO implement me
	panic("implement me")
}
