package client

type EventOverData struct {
	GamelogURL    string `json:"gamelogURL"`
	VisualizerURL string `json:"visualizerURL"`
	Message       string `json:"message"`
}

func WaitForEventOver() EventOverData {
	data := EventOverData{}
	WaitForEvent("lobbied", &data)

	return data
}
