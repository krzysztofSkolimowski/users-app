package adapters

import (
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

	// todo - handle unique user id constraint
	exists, err := r.db.Collection("users").Find(db.Cond{"id": user.ID}).Exists()
	if err != nil {
		return err
	}

	if exists {
		return domain.ErrUserAlreadyExists
	}

	return r.insert(user)
}

func (r repository) UpdateUser(user domain.User) error {
	exists, err := r.db.Collection("users").Find(db.Cond{"id": user.ID}).Exists()
	if err != nil {
		return err
	}

	if !exists {
		return domain.ErrUserNotFound
	}

	return r.db.Collection("users").Find(db.Cond{"id": user.ID}).Update(fromDomain(user))
}

func (r repository) RemoveUser(id domain.UserID) error {
	res := r.db.Collection("users").Find(db.Cond{"id": id})
	return res.Delete()
}

func (r repository) Users(filter string) ([]domain.User, error) {
	//TODO implement me
	panic("implement me")
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
