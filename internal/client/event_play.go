package client

// EventPlayData is the expected shape of the event "play"'s data.
type EventPlayData struct {
	ClientType       string `json:"clientType"`
	GameName         string `json:"gameName"`
	GameSettings     string `json:"gameSettings"`
	Password         string `json:"password"`
	PlayerIndex      int    `json:"playerIndex"`
	PlayerName       string `json:"playerName"`
	RequestedSession string `json:"requestedSession"`
}

// SendEventPlay sends the event "play" to the game server.
func SendEventPlay(data EventPlayData) {
	sendEvent("play", data)
}
