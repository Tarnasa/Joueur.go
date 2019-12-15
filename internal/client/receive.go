package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"joueur/internal/errorhandler"

	"github.com/fatih/color"
)

// BaseEvent is the base shape of an event being sent or received from the
// Cadre game server.
type BaseEvent struct {
	EventName string      `json:"event"`
	Data      interface{} `json:"data"`
}

const readSize = 1024

var eventsStack = make([][]byte, 0)
var receivedBuffer = make([]byte, 0)

// waitForEvents will block the call stack and wait for the TCP connection to
// fill with at least one event. Once 1 or more bytes of an event are received
// this function returns assuming something will parse those bytes.
func waitForEvents() {
	if len(eventsStack) > 0 {
		return
	}

	for {
		sent := make([]byte, readSize)
		bytesSent, err := (*instance.conn).Read(sent)
		if err != nil {
			errorhandler.HandleError(
				errorhandler.CannotReadSocket,
				err,
				"Error reading socket while waiting for events",
			)
		}

		if bytesSent == 0 {
			continue
		}

		sent = sent[:bytesSent] // cut off bytes not sent as they are junk 0's

		if instance.printIO {
			color.Magenta("FROM SERVER --> " + string(sent))
		}

		byteArrays := bytes.Split(sent, []byte{eotChar})
		numberOfByteArrays := len(byteArrays)
		// the first element should be a continuation of what we last read
		byteArrays[0] = append(receivedBuffer, byteArrays[0]...)
		if numberOfByteArrays == 1 {
			// then the entire read had no EOT chars,
			// so just buffer more bytes
			receivedBuffer = byteArrays[0]
			continue
		}
		// else with at least 2 elements split we received the EOT character
		// so we have at least one event to parse from the bytes.
		receivedBuffer = byteArrays[numberOfByteArrays-1] // buffer the last slice
		byteArrays = byteArrays[:numberOfByteArrays-1]    // and remove it

		//
		for _, eventBytes := range byteArrays {
			eventsStack = append(eventsStack, eventBytes)
		}

		if len(eventsStack) > 0 {
			return
		}
	}
}

// waitForEvent waits for a given event name by its string and given a
// destination interface stores the resulting data shape in that destination.
func waitForEvent(eventName string, dataDestination interface{}) {
	for {
		waitForEvents()

		for len(eventsStack) > 0 {
			// pop first event off the front of the events stack
			eventBytes := eventsStack[0]
			eventsStack = eventsStack[1:]
			var baseEvent *BaseEvent = nil
			nameErr := json.Unmarshal(eventBytes, &baseEvent)

			if baseEvent == nil {
				nameErr = errors.New("No parsed base event")
			}

			if nameErr != nil {
				errorhandler.HandleError(
					errorhandler.MalformedJSON,
					nameErr,
					"Could not parse base JSON"+string(eventBytes),
				)
			}

			if baseEvent.EventName == eventName {
				destination := &BaseEvent{
					EventName: eventName,
					Data:      dataDestination,
				}

				err := json.Unmarshal(eventBytes, destination)

				if dataDestination == nil {
					err = errors.New("No destination data to unmarshal data into")
				}

				if err != nil {
					errorhandler.HandleError(
						errorhandler.MalformedJSON,
						err,
						"Error occurred while waiting for "+eventName,
					)
				}

				return
			}
			// else attempt to auto handle the event
			switch baseEvent.EventName {
			case "delta":
				autoHandleEventDelta(eventBytes)
			case "fatal":
				autoHandleEventFatal(eventBytes)
			case "order":
				autoHandleEventOrder(eventBytes)
			case "invalid":
				autoHandleEventInvalid(eventBytes)
			default:
				errorhandler.HandleError(
					errorhandler.UnknownEventFromServer,
					errors.New("No event auto handler for "+baseEvent.EventName),
					"Unknown event could not be handled",
				)
			}
		}
	}
}
