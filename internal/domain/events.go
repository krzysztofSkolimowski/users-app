package domain

import "context"

// List of possible and currently supported events
var (
	UserAdded    = EventMsg("user-added")
	UserModified = EventMsg("user-modified")
	UserDeleted  = EventMsg("user-deleted")
)

// EventMsg is used to identify the type of event
type EventMsg string

// Event represents an application event, which includes the command that triggered the event and the event message.
type Event struct {
	Command
	Msg EventMsg
}

type Publisher interface {
	PublishEvent(context.Context, Event) error
}

// Command is an interface that defines the EncodeEvent method, which is responsible for encoding a command into an event.
type Command interface {
	EncodeEvent() (Event, error)
}
