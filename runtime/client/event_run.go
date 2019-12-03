package client

type GameObjectReference struct {
	Id string `json:"id"`
}

type EventRun struct {
	Caller       GameObjectReference    `json:"caller"`
	FunctionName string                 `json:"functionName"`
	Args         map[string]interface{} `json:"args"`
}

func SendEventRun(data EventRun) {
	SendEvent("run", data)
}
