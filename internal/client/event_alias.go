package client

// SendEventAlias sends to the server the "alias" event
func SendEventAlias(gameName string) {
	sendEvent("alias", gameName)
}
