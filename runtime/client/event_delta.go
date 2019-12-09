package client

import (
	"encoding/json"
	"errors"
	"joueur/runtime/errorhandler"
)

type eventDelta struct {
	Data map[string]interface{} `json:"data"`
}

// EventDeltaHandler is the handler that should be set by the GameManager to
// handler "delta" events, which cases a delta merge.
var EventDeltaHandler func(map[string]interface{}) = nil

func autoHandleEventDelta(eventBytes []byte) {
	var parsed eventDelta
	json.Unmarshal(eventBytes, &parsed)

	if EventDeltaHandler == nil {
		errorhandler.HandleError(
			errorhandler.DeltaMergeFailure,
			errors.New("cannot merge delta without EventDeltaHandler set"),
		)
	}
	EventDeltaHandler(parsed.Data)
}
