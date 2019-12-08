package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"joueur/runtime/errorhandler"
)

type EventOrderData struct {
	Name  string        `json:"name"`
	Index int64         `json:"index"`
	Args  []interface{} `json:"args"`
}

type eventOrder struct {
	Data EventOrderData `json:"data"`
}

var EventOrderHandler func(e EventOrderData) = nil

func autoHandleEventOrder(eventBytes []byte) {
	fmt.Println("auto handling order...", string(eventBytes))

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
