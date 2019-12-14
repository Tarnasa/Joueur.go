package client

// ServerConstants are the expected delta merge constants the game server
// will be using.
type ServerConstants struct {
	DeltaRemoved       string `json:"DELTA_REMOVED"`
	DeltaListLengthKey string `json:"DELTA_LIST_LENGTH"`
}

// EventLobbiedData is the shape of the "lobbied" data event.
type EventLobbiedData struct {
	// The name of the game you are playing
	GameName string `json:"gameName"`
	// The version of the game being ran on the server.
	GameVersion string `json:"gameVersion"`
	// The game session (id) of the game you will be playing.
	GameSession string `json:"gameSession"`
	// Constants used to facilitate game IO communication.
	Constants ServerConstants `json:"constants"`
}

// WaitForEventLobbied waits for the "lobbied" event and returns the resulting
// data from that event.
func WaitForEventLobbied() EventLobbiedData {
	data := EventLobbiedData{}
	waitForEvent("lobbied", &data)

	return data
}
