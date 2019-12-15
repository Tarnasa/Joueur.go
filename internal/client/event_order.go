package client

import (
	"encoding/json"
	"errors"
	"joueur/internal/errorhandler"
)

// EventOrderData is the expected shape of an "order" event data
type EventOrderData struct {
	Name  string        `json:"name"`
	Index int64         `json:"index"`
	Args  []interface{} `json:"args"`
}

type eventOrder struct {
	Data EventOrderData `json:"data"`
}

// EventOrderHandler is the handler for "order" events. GameManager should
// inject into this.
var EventOrderHandler func(e EventOrderData) = nil

// autoHandleEventOrder automatically handles any "order" events by parsing it
// and invoking the EventOrderHandler.
func autoHandleEventOrder(eventBytes []byte) {
	var parsed eventOrder

	err := json.Unmarshal(eventBytes, &parsed)
	if err != nil {
		errorhandler.HandleError(
			errorhandler.MalformedJSON,
			err,
			"Could not parse order event",
		)
	}

	if EventOrderHandler == nil {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("no event order auto handler in client"),
		)
	}
	EventOrderHandler(parsed.Data)
}
