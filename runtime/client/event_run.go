package client

// GameObjectReference is a container for a game object reference with just
// its id.
type GameObjectReference struct {
	ID string `json:"id"`
}

// EventRunData is the data shape expected for a "run" event.
type EventRunData struct {
	Caller       GameObjectReference    `json:"caller"`
	FunctionName string                 `json:"functionName"`
	Args         map[string]interface{} `json:"args"`
}

// SendEventRun sends the "run" event to the game server.
func SendEventRun(data EventRunData) {
	sendEvent("run", data)
}
