package service

import (
	"context"
	"fmt"
	"users-app/domain"

	"go.uber.org/zap"
)

// CommandLoggingWrapper is a wrapper for UsersCommandService
// It logs the commands received and the errors returned
// It's purpose is to decouple the application from the logger
type CommandLoggingWrapper struct {
	wrapped UsersCommandService
	logger  Logger
}

func NewCommandLoggingWrapper(logger Logger, wrapped UsersCommandService) UsersCommandService {
	return CommandLoggingWrapper{wrapped, logger}
}

func (c CommandLoggingWrapper) AddUser(ctx context.Context, command AddUserCommand) (domain.User, error) {
	c.logger.Info(fmt.Sprintf("AddUser command received: %v", command))
	u, err := c.wrapped.AddUser(ctx, command)
	if err != nil {
		c.logger.Error(fmt.Sprintf("AddUser command failed: %v", err))
		return domain.User{}, err
	}

	return u, nil
}

func (c CommandLoggingWrapper) ModifyUser(ctx context.Context, command ModifyUserCommand) error {
	c.logger.Info(fmt.Sprintf("ModifyUser command received: %v", command))
	err := c.wrapped.ModifyUser(ctx, command)
	if err != nil {
		c.logger.Error(fmt.Sprintf("ModifyUser command failed: %v", err))
		return err
	}

	return nil
}

func (c CommandLoggingWrapper) DeleteUser(ctx context.Context, command DeleteUserCommand) error {
	c.logger.Info(fmt.Sprintf("DeleteUser command received: %v", command))
	err := c.wrapped.DeleteUser(ctx, command)
	if err != nil {
		c.logger.Error(fmt.Sprintf("DeleteUser command failed: %v", err))
		return err
	}

	return nil

}

type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
}
