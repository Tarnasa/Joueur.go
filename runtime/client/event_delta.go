package client

import (
	"encoding/json"
)

type EventDelta struct {
	event string
	data map[string]interface{}
}

func autoHandleEventDelta(eventBytes []byte) {
	deltaEvent := EventDelta{
		event: "delta",
		data: make(map[string]interface{}),
	}
	json.Unmarshal(eventBytes, &deltaEvent)

	eventDeltaHandler(deltaEvent.data)
}
