package client

type EventFinishedData struct {
	OrderIndex int64       `json:"orderIndex"`
	Returned   interface{} `json:"returned"`
}

func SendEventFinished(data EventFinishedData) {
	SendEvent("finished", data)
}
