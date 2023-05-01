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

func ParseID(id string) (UserID, error) {
	return uuid.Parse(id)
}

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

type Field string

// Fields represent user fields that can be changes
type Fields map[Field]string

type Filter struct {
	FirstName *string
	LastName  *string
	Nickname  *string
	Email     *string
	Country   *string
}

var MaxPaginationLimit = 100
var DefaultPaginationLimit = 10

type Pagination struct {
	limit  int
	Offset int
}

func (p Pagination) Limit() int {
	return p.limit
}

func NewPagination(limit, offset int) Pagination {
	ret := DefaultPaginationLimit
	if limit != 0 {
		ret = max(limit, MaxPaginationLimit)
	}

	return Pagination{
		limit:  ret,
		Offset: offset,
	}
}

func max(a, b int) int {
	if a < b {
		return a
	}

	return b
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

type Repository interface {
	AddUser(User) error
	ModifyUser(UserID, Fields) error
	RemoveUser(UserID) error
	Users(Filter, Pagination) ([]User, error)
}
