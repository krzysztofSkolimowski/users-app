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

// User is a domain entity representing a single user
// I am assuming that password should not be returned from the repository
// since it's a security risk
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

// Field represents a single user field that can be changed
type Field string

// Fields represent user fields that can be changes
type Fields map[Field]string

func (f Fields) AddIfValIsNotNil(key Field, value *string) {
	if value != nil {
		f[key] = *value
	}
}

func (f Filter) FirstName() *string { return f.firstName }
func (f Filter) LastName() *string  { return f.lastName }
func (f Filter) Nickname() *string  { return f.nickname }
func (f Filter) Email() *string     { return f.email }
func (f Filter) Country() *string   { return f.country }

// Filter represents a filter that can be used to search users
// In case a field is nil, it will not be used in the search
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

func valueIfNotEmpty(value string) *string {
	if value == "" {
		return nil
	}

	return &value
}

func NewFilter(firstName, lastName, nickname, email, country string) Filter {
	filter := Filter{}

	filter.firstName = valueIfNotEmpty(firstName)
	filter.lastName = valueIfNotEmpty(lastName)
	filter.nickname = valueIfNotEmpty(nickname)
	filter.email = valueIfNotEmpty(email)
	filter.country = valueIfNotEmpty(country)

	return filter
}

var MaxPaginationLimit = 100
var DefaultPaginationLimit = 10
var DefaultPagination = Pagination{DefaultPaginationLimit, 0}

func (p Pagination) Limit() int { return p.limit }

type Pagination struct {
	limit  int
	Offset int
}

func NewPagination(limit, offset int) Pagination {
	if offset < 0 || limit < 0 {
		return DefaultPagination
	}

	ret := DefaultPaginationLimit
	if limit != 0 {
		ret = min(limit, MaxPaginationLimit)
	}

	return Pagination{
		limit:  ret,
		Offset: offset,
	}
}

func min(a, b int) int {
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
