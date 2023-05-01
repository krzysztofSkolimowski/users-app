package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
)

type UserID = uuid.UUID

type User struct {
	ID           UserID
	FirstName    string
	LastName     string
	Nickname     string
	PasswordHash string
	Email        string
	Country      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewUser(
	firstName string, lastName string, nickname string,
	password string, email string, country string,
) User {
	return User{
		ID:           uuid.New(),
		FirstName:    firstName,
		LastName:     lastName,
		Nickname:     nickname,
		PasswordHash: hash(password),
		Email:        email,
		Country:      country,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func hash(password string) string {
	// todo - implement proper hashing with salt
	// Hash the password using SHA-256

	hashed := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hashed[:])
}

// todo - add context
type Repository interface {
	AddUser(user User) error
	UpdateUser(user User) error
	RemoveUser(id UserID) error
	Users(filter string) ([]User, error)
}
