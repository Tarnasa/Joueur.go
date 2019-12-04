package client

type GameObjectReference struct {
	Id string `json:"id"`
}

type EventRunData struct {
	Caller       GameObjectReference    `json:"caller"`
	FunctionName string                 `json:"functionName"`
	Args         map[string]interface{} `json:"args"`
}

func SendEventRun(data EventRunData) {
	SendEvent("run", data)
}
