package client

type ServerConstants struct {
	DeltaRemoved       string `json:"DELTA_REMOVED"`
	DeltaListLengthKey string `json:"DELTA_LIST_LENGTH"`
}

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

func WaitForEventLobbied() EventLobbiedData {
	data := EventLobbiedData{}
	WaitForEvent("lobbied", &data)

	return data
}
