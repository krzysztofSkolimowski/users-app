package domain

import "time"

type User struct {
	ID        string
	FirstName string
	LastName  string
	Nickname  string
	Password  string
	Email     string
	Country   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(
	id string, firstName string, lastName string, nickname string,
	password string, email string, country string,
) User {
	return User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Nickname:  nickname,
		Password:  password,
		Email:     email,
		Country:   country,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type Repository interface {
	AddUser(user User) error
	UpdateUser(user User) error
	RemoveUser(id string) error
	Users(filter string) ([]User, error)
}
