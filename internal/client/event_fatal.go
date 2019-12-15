package client

import (
	"errors"
	"joueur/internal/errorhandler"
)

// autoHandleEventFatal automatically handles "fatal" events by reporting
// the error.
func autoHandleEventFatal(eventBytes []byte) {
	errorhandler.HandleError(
		errorhandler.FatalEvent,
		errors.New("Unexpected fatal event from server"),
		"Got a fatal event from the server: "+string(eventBytes),
	)
}
