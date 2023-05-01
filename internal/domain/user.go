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
	ErrEmailExists       = errors.New("email already exists")
	ErrEmailRequired     = errors.New("email is required")
)

type UserID = uuid.UUID

func NewUserID() UserID {
	return uuid.New()
}

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

func (f Filter) FirstName() *string { return f.firstName }
func (f Filter) LastName() *string  { return f.lastName }
func (f Filter) Nickname() *string  { return f.nickname }
func (f Filter) Email() *string     { return f.email }
func (f Filter) Country() *string   { return f.country }

type Filter struct {
	firstName *string
	lastName  *string
	nickname  *string
	email     *string
	country   *string
}

func NewFilterEmail(email string) Filter {
	return Filter{email: &email}
}

func NewFilter(firstName, lastName, nickname, email, country string) Filter {
	filter := Filter{}
	if firstName != "" {
		filter.firstName = &firstName
	}
	if lastName != "" {
		filter.lastName = &lastName
	}
	if nickname != "" {
		filter.nickname = &nickname
	}
	if email != "" {
		filter.email = &email
	}
	if country != "" {
		filter.country = &country
	}

	return filter
}

var MaxPaginationLimit = 100
var DefaultPaginationLimit = 10
var DefaultPagination = Pagination{DefaultPaginationLimit, 0}

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
) (User, error) {
	if email == "" {
		return User{}, ErrEmailRequired
	}

	return User{
		ID:           NewUserID(),
		FirstName:    firstName,
		LastName:     lastName,
		Nickname:     nickname,
		PasswordHash: hash(password),
		Email:        email,
		Country:      country,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}, nil
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
