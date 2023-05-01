package domain

import "context"

type Event struct {
	Command
	Msg string
}

type Publisher interface {
	PublishEvent(context.Context, Event) error
}

type Command interface {
	EncodeEvent() (Event, error)
}
