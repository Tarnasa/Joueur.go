package client

type EventStartData struct {
	PlayerID string
}

func WaitForEventStart() EventStartData {
	data := EventStartData{}
	WaitForEvent("start", &data)

	return data
}
