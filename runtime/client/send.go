package client

import (
	"encoding/json"
	"errors"
	"joueur/runtime/errorhandler"
	"time"

	"github.com/fatih/color"
)

// sendEventData is the basic shape of an event we send to the server.
type sendEventData struct {
	Event    string      `json:"event"`
	SentTime int64       `json:"sentTime"`
	Data     interface{} `json:"data"`
}

// sendEvent sends an even by its name with some optional data to the game
// server.
func sendEvent(event string, data interface{}) {
	bytes, err := json.Marshal(sendEventData{
		Event:    event,
		Data:     data,
		SentTime: time.Now().Unix(),
	})

	if err != nil {
		errorhandler.HandleError(
			errorhandler.MalformedJSON,
			err,
			"Could not encode event to json",
		)
	}

	sendRaw(append(bytes, eotChar))
}

// sendRaw sends raw bytes to the game server. It's assumed these bytes are in
// the correct serialized event format.
func sendRaw(bytes []byte) error {
	/**
	 * Sends a raw string to the game server
	 * @param str The string to send.
	 */
	if instance.conn == nil {
		return errors.New("Cannot write to socket that has not been initialized")
	}

	if instance.printIO {
		color.Magenta("TO SERVER <-- " + string(bytes))
	}

	_, err := (*instance.conn).Write((bytes))
	if err != nil {
		errorhandler.HandleError(
			errorhandler.DisconnectedUnexpectedly,
			err,
			"Could not send string through server.",
		)
	}

	return nil
}
