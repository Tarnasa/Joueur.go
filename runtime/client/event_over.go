package client

// EventOverData is the shape of the data object in an "over" event.
type EventOverData struct {
	GamelogURL    string `json:"gamelogURL"`
	VisualizerURL string `json:"visualizerURL"`
	Message       string `json:"message"`
}

// WaitForEventOver will wait for the server to send the "over" event.
// This is expected to take a LONG time, ideally until the game is actually
// over, which wil; end this program.
func WaitForEventOver() EventOverData {
	data := EventOverData{}
	// we expect this to take a LONG time...
	WaitForEvent("over", &data)

	return data
}
