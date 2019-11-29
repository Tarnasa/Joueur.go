package client

type EventPlay struct {
	ClientType       string `json:"clientType"`
	GameName         string `json:"gameName"`
	GameSettings     string `json:"gameSettings"`
	Password         string `json:"password"`
	PlayerIndex      int    `json:"playerIndex"`
	PlayerName       string `json:"playerName"`
	RequestedSession string `json:"requestedSession"`
}

func SendEventPlay(data EventPlay) {
	SendEvent("play", data)
}
