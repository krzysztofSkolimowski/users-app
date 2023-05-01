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

func NewRepository() domain.Repository {
	// todo - load from .env
	settings := postgresql.ConnectionURL{
		Host:     "db",
		Database: "users",
		User:     "postgres",
		Password: "password",
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
	_, err := r.db.Collection("users").Insert(fromDomain(user))
	return err
}

func (r repository) UpdateUser(user domain.User) error {
	res := r.db.Collection("users").Find(db.Cond{"id": user.ID})

	return res.Update(fromDomain(user))
}

func (r repository) RemoveUser(id domain.UserID) error {
	res := r.db.Collection("users").Find(db.Cond{"id": id})
	return res.Delete()
}

func (r repository) Users(filter string) ([]domain.User, error) {
	//TODO implement me
	panic("implement me")
}
