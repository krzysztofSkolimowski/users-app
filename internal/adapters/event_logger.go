package adapters

import (
	"log"
	"os"
	"users-app/domain"
)

// EventLogger is a specialized logger for logging events and storing them in the logs/event.log file.
// The application is event-driven, so maintaining a correct event log is essential for replaying all events
// since the application's startup to recreate its state at a given point in time.
//
// The EventLogger must be closed using its Close method when it's no longer needed, to ensure that the underlying
// file is properly closed and resources are released.
type EventLogger struct {
	logger *log.Logger
	file   *os.File
}

func NewEventLogger(filename string) *EventLogger {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file, "", log.LstdFlags)
	return &EventLogger{
		logger: logger,
		file:   file,
	}
}

// LogEvent logs an event to the file
func (e *EventLogger) LogEvent(event domain.Event) {
	e.logger.Printf("%#v", event)
}

func (e *EventLogger) Close() error {
	return e.file.Close()
}
