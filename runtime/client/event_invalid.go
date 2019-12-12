package client

import (
	"encoding/json"
	"errors"
	"joueur/runtime/errorhandler"
)

// EventInvalidData is the shape of the "invalid" data event.
type eventInvalidData struct {
	// The human readable reason why something invalid happened
	Data struct {
		Message string `json:"message"`
	} `json:"data"`
}

// EventInvalidHandler is the handler to automatically handle "invalid" events
var EventInvalidHandler func(string)

// autoHandleEventOrder automatically handles "invalid" events
func autoHandleEventInvalid(eventBytes []byte) {
	var parsed eventInvalidData

	err := json.Unmarshal(eventBytes, &parsed)
	if err != nil {
		errorhandler.HandleError(
			errorhandler.MalformedJSON,
			err,
			"Could not parse invalid event",
		)
	}

	if EventInvalidHandler == nil {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("no invalid event auto handler in client"),
		)
	}
	EventInvalidHandler(parsed.Data.Message)
}
