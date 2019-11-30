package client

import (
	"encoding/json"
	"fmt"
)

type EventDelta struct {
	Data map[string]interface{} `json:"data"`
}

func autoHandleEventDelta(eventBytes []byte) {
	fmt.Println("auto hanlding delta...")

	var parsed EventDelta

	err := json.Unmarshal(eventBytes, &parsed)
	fmt.Println("auto handled the delta and got...", parsed, string(eventBytes), err)

	eventDeltaHandler(parsed.Data)
}
