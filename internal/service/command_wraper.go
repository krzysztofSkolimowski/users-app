package service

import (
	"context"
	"log"
	"users-app/domain"
)

type CommandEventsWrapper struct {
	publisher domain.Publisher
	wrapped   UsersCommandService
}

func NewCommandEventsWrapper(publisher domain.Publisher, wrapped UsersCommandService) UsersCommandService {
	return CommandEventsWrapper{publisher: publisher, wrapped: wrapped}
}

func (c AddUserCommand) EncodeEvent() (domain.Event, error) {
	return domain.Event{Msg: "user-added", Command: c}, nil
}

func (c CommandEventsWrapper) AddUser(ctx context.Context, command AddUserCommand) (domain.User, error) {
	go c.publishEvent(ctx, command)
	return c.wrapped.AddUser(ctx, command)
}

func (c ModifyUserCommand) EncodeEvent() (domain.Event, error) {
	return domain.Event{Msg: "user-modified", Command: c}, nil
}

func (c CommandEventsWrapper) ModifyUser(ctx context.Context, command ModifyUserCommand) error {
	go c.publishEvent(ctx, command)
	return c.wrapped.ModifyUser(ctx, command)
}

func (c DeleteUserCommand) EncodeEvent() (domain.Event, error) {
	return domain.Event{Msg: "user-deleted", Command: c}, nil
}

func (c CommandEventsWrapper) DeleteUser(ctx context.Context, command DeleteUserCommand) error {
	go c.publishEvent(ctx, command)
	return c.wrapped.DeleteUser(ctx, command)
}

// publishEvent is a helper function that publishes an event to the publisher
// I am consciously choosing to ignore errors here, since I don't want to block the execution of the command
// if the event fails to be published.
// This is a tradeoff, in the next version of this code I would run a publisher in a separate goroutine,
// add buffer for the events and implement a retry mechanism to make sure no events are lost
func (c CommandEventsWrapper) publishEvent(ctx context.Context, command domain.Command) {
	event, err := command.EncodeEvent()
	if err != nil {
		log.Printf("error publishing event: %v", err)
	}

	err = c.publisher.PublishEvent(ctx, event)
	if err != nil {
		log.Printf("error publishing event: %v", err)
	}

	return
}
