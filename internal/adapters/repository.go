package adapters

import (
	"fmt"
	"log"
	"users-app/domain"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

type repository struct {
	db db.Session
}

type RepoConfig struct {
	Host     string
	Database string
	User     string
	Password string
}

func NewRepository(
	repositoryConfig RepoConfig,
) repository {
	// todo - load from .env
	settings := postgresql.ConnectionURL{
		Host:     repositoryConfig.Host,
		Database: repositoryConfig.Database,
		User:     repositoryConfig.User,
		Password: repositoryConfig.Password,
	}

	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal(err)
	}

	err = sess.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return repository{db: sess}
}

func (r repository) AddUser(user domain.User) error {
	exists, err := r.db.Collection("users").Find(db.Cond{"id": user.ID}).Exists()
	if err != nil {
		return err
	}
	if exists {
		return domain.ErrUserAlreadyExists
	}

	return r.insert(user)
}

func (r repository) ModifyUser(id domain.UserID, fields domain.Fields) error {
	res := r.db.Collection("users").Find(db.Cond{"id": id})
	exists, err := res.Exists()
	if err != nil {
		return err
	}
	if !exists {
		return domain.ErrUserNotFound
	}

	if len(fields) > 0 {
		err := res.Update(fields)
		if err != nil {
			return fmt.Errorf("failed to update user: %w", err)
		}

		return nil
	}

	return nil
}

func (r repository) RemoveUser(id domain.UserID) error {
	res := r.db.Collection("users").Find(db.Cond{"id": id})
	return res.Delete()
}

func (r repository) Users(filter domain.Filter, pagination domain.Pagination) ([]domain.User, error) {
	query := r.db.Collection("users").Find()
	query = addFilters(filter, query)

	// add pagination
	query = query.Limit(pagination.Limit())
	query = query.Offset(pagination.Offset)

	// Execute the query and fetch the results
	var ret []UserDTO
	err := query.All(&ret)
	if err != nil {
		return nil, err
	}

	return toDomainUsers(ret), nil
}

func addFilters(filter domain.Filter, q db.Result) db.Result {
	query := q
	filterMap := map[string]*string{
		"first_name": filter.FirstName,
		"last_name":  filter.LastName,
		"nickname":   filter.Nickname,
		"email":      filter.Email,
		"country":    filter.Country,
	}

	for field, value := range filterMap {
		if value != nil {
			query = query.And(db.Cond{fmt.Sprintf("%s", field): *value})
		}
	}

	return query
}

func (r repository) insert(u domain.User) error {
	_, err := r.db.Collection("users").Insert(fromDomain(u))
	return err
}

// implemented just for integration tests
func (r repository) flush() {
	r.db.Collection("users").Truncate()
}

// implemented just for integration tests
func (r repository) allUsers() ([]domain.User, error) {
	var users []UserDTO
	err := r.db.Collection("users").Find().All(&users)
	return toDomainUsers(users), err
}
