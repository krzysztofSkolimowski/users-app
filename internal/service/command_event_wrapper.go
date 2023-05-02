package service

import (
	"context"
	"log"
	"users-app/domain"
)

// CommandEventsWrapper is a wrapper around UsersCommandService that logs and publishes events based on commands.
// It is used to decouple the application from the event publisher and logger.
// As application uses CQRS patter, the Commands are the only way to change the state of the application.
// Therefore, every command triggers an event - which symbolizes a change in the application's state.
// The event is then published to the Redis channel and logged to the logs/event.log file.
type CommandEventsWrapper struct {
	publisher   domain.Publisher
	wrapped     UsersCommandService
	eventLogger eventLogger
}

func NewCommandEventsWrapper(publisher domain.Publisher, wrapped UsersCommandService, logger eventLogger) UsersCommandService {
	return CommandEventsWrapper{publisher, wrapped, logger}
}

// eventLogger is a specialized logger for logging events and storing them in the logs/event.log file.
// The application is event-driven, so maintaining a correct event log is essential for replaying all events
// since the application's startup to recreate its state at a given point in time.
type eventLogger interface {
	LogEvent(domain.Event)
}

func (c AddUserCommand) EncodeEvent() (domain.Event, error) {
	return domain.Event{Msg: domain.UserAdded, Command: c}, nil
}

func (c CommandEventsWrapper) AddUser(ctx context.Context, command AddUserCommand) (domain.User, error) {
	c.publishEvent(ctx, command)
	return c.wrapped.AddUser(ctx, command)
}

func (c ModifyUserCommand) EncodeEvent() (domain.Event, error) {
	return domain.Event{Msg: domain.UserModified, Command: c}, nil
}

func (c CommandEventsWrapper) ModifyUser(ctx context.Context, command ModifyUserCommand) error {
	c.publishEvent(ctx, command)
	return c.wrapped.ModifyUser(ctx, command)
}

func (c DeleteUserCommand) EncodeEvent() (domain.Event, error) {
	return domain.Event{Msg: domain.UserDeleted, Command: c}, nil
}

func (c CommandEventsWrapper) DeleteUser(ctx context.Context, command DeleteUserCommand) error {
	c.publishEvent(ctx, command)
	return c.wrapped.DeleteUser(ctx, command)
}

// publishEvent is a helper function that publishes an event to the publisher without blocking the execution of the command.
// It logs the event and publishes it asynchronously. Errors during publishing are logged but do not block the command execution.
// This is a tradeoff, and future versions of the code may implement a separate goroutine, event buffering, and a retry mechanism.
func (c CommandEventsWrapper) publishEvent(ctx context.Context, command domain.Command) {
	event, err := command.EncodeEvent()
	if err != nil {
		log.Printf("error publishing event: %v", err)
	}
	c.eventLogger.LogEvent(event)

	go func() {
		err = c.publisher.PublishEvent(ctx, event)
		if err != nil {
			log.Printf("error publishing event: %v", err)
		}
	}()

	return
}
