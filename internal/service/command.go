package service

import (
	"context"
	"users-app/domain"
)

type UsersCommandService interface {
	AddUser(context.Context, AddUserCommand) (domain.User, error)
	ModifyUser(context.Context, ModifyUserCommand) error
	DeleteUser(context.Context, DeleteUserCommand) error
}

type UserCommandService struct {
	userRepository domain.Repository
}

func NewUserCommandService(repo domain.Repository) *UserCommandService {
	return &UserCommandService{repo}
}

// AddUserCommand is used to add a new user
type AddUserCommand struct {
	FirstName string
	LastName  string
	Nickname  string
	Password  string
	Email     string
	Country   string
}

func (u UserCommandService) AddUser(ctx context.Context, toAdd AddUserCommand) (domain.User, error) {
	user, err := domain.NewUser(toAdd.FirstName, toAdd.LastName, toAdd.Nickname, toAdd.Password, toAdd.Email, toAdd.Country)
	if err != nil {
		return domain.User{}, err
	}

	users, err := u.userRepository.Users(domain.NewFilterEmail(toAdd.Email), domain.DefaultPagination)
	if err != nil {
		// todo - handle errors
		return domain.User{}, err
	}

	if len(users) > 0 {
		return domain.User{}, domain.ErrEmailExists
	}

	err = u.userRepository.AddUser(user)
	if err != nil {
		// todo - handle errors
		return domain.User{}, err
	}

	return user, nil
}

// ModifyUserCommand is used to modify user data
// I am assuming that password modifications should not be possible from this command,
// since it's a security risk and it would probably require some additional checks
type ModifyUserCommand struct {
	ID        domain.UserID
	FirstName *string
	LastName  *string
	Nickname  *string
	Email     *string
	Country   *string
}

func (c ModifyUserCommand) fieldsToUpdate() domain.Fields {
	ret := make(domain.Fields)
	if c.FirstName != nil {
		ret["first_name"] = *c.FirstName
	}
	if c.LastName != nil {
		ret["last_name"] = *c.LastName
	}
	if c.Nickname != nil {
		ret["nickname"] = *c.Nickname
	}
	if c.Email != nil {
		ret["email"] = *c.Email
	}
	if c.Country != nil {
		ret["country"] = *c.Country
	}

	return ret
}

func (u UserCommandService) ModifyUser(ctx context.Context, toModify ModifyUserCommand) error {
	err := u.userRepository.ModifyUser(toModify.ID, toModify.fieldsToUpdate())
	if err != nil {
		return err
	}

	return nil
}

// DeleteUserCommand is used to delete a user given the user exists
type DeleteUserCommand struct {
	ID domain.UserID
}

func (u UserCommandService) DeleteUser(ctx context.Context, toDelete DeleteUserCommand) error {
	return u.userRepository.RemoveUser(toDelete.ID)
}
