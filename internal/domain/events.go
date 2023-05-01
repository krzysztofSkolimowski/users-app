package domain

import "context"

var (
	UserAdded    = EventMsg("user-added")
	UserModified = EventMsg("user-modified")
	UserDeleted  = EventMsg("user-deleted")
)

type EventMsg string

type Event struct {
	Command
	Msg EventMsg
}

type Publisher interface {
	PublishEvent(context.Context, Event) error
}

type Command interface {
	EncodeEvent() (Event, error)
}
