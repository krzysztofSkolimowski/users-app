package service

import (
	"context"
	"users-app/domain"
)

type UsersCommandService interface {
	AddUser(context.Context, AddUserCommand) (domain.User, error)
	ModifyUser(context.Context, ModifyUserCommand) (domain.User, error)
	DeleteUser(context.Context, DeleteUserCommand) error
}

type UserCommandService struct {
	userRepository domain.Repository
}

func NewUserCommandService(repo domain.Repository) *UserCommandService {

	// todo - implement normal repo
	// mock repo

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
	user := domain.NewUser(toAdd.FirstName, toAdd.LastName, toAdd.Nickname, toAdd.Password, toAdd.Email, toAdd.Country)

	// todo - check if user exists
	err := u.userRepository.AddUser(user)
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
	FirstName string
	LastName  string
	Nickname  string
	Email     string
	Country   string
}

func (u UserCommandService) ModifyUser(ctx context.Context, toModify ModifyUserCommand) (domain.User, error) {
	// todo - probably implement patch update
	user := domain.User{
		ID:        toModify.ID,
		FirstName: toModify.FirstName,
		LastName:  toModify.LastName,
		Nickname:  toModify.Nickname,
		Email:     toModify.Email,
		Country:   toModify.Country,
	}

	err := u.userRepository.UpdateUser(user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

// DeleteUserCommand is used to delete a user given the user exists
type DeleteUserCommand struct {
	ID domain.UserID
}

func (u UserCommandService) DeleteUser(ctx context.Context, toDelete DeleteUserCommand) error {
	return u.userRepository.RemoveUser(toDelete.ID)
}
