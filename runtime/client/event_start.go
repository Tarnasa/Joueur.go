package client

// EventStartData is the expected shape of the data from a "start" event.
type EventStartData struct {
	PlayerID string
}

// WaitForEventStart waits for the "start" event from the game server and
// returns the event's data.
func WaitForEventStart() EventStartData {
	data := EventStartData{}
	waitForEvent("start", &data)

	return data
}
