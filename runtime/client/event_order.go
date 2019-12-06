package client

import (
	"encoding/json"
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

var EventOverHandler func(e EventOrderData)

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

	EventOverHandler(parsed.Data)
}
