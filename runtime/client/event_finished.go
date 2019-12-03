package client

type EventFinishedData struct {
	OrderIndex int64       `json:"orderIndex"`
	returned   interface{} `json:"returned"`
}

func SendEventFinished(data EventFinishedData) {
	SendEvent("finished", data)
}
