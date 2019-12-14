package client

// EventFinishedData is the expected shape of data from the "finished" event
type EventFinishedData struct {
	OrderIndex int64       `json:"orderIndex"`
	Returned   interface{} `json:"returned"`
}

// SendEventFinished sends the "finished" event to the game server
func SendEventFinished(data EventFinishedData) {
	sendEvent("finished", data)
}
