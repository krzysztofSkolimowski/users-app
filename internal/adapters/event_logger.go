package adapters

import (
	"log"
	"os"
	"users-app/domain"
)

type EventLogger struct {
	logger *log.Logger
}

func NewEventLogger(filename string) *EventLogger {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)
	return &EventLogger{
		logger: logger,
	}
}

func (e *EventLogger) LogEvent(event domain.Event) {
	e.logger.Printf("%#v", event)
}
