package client

type EventNamed struct {
	Event string `json:"event"`
	Data string `json:"data"`
}

func WaitForEventNamed() string {
	var named EventNamed
	WaitForEvent("named", &named)

	return named.Data
}
