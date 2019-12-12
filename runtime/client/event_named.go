package client

// WaitForEventNamed waits for the "named" event and returns that string.
func WaitForEventNamed() string {
	named := ""
	waitForEvent("named", &named)

	return named
}
