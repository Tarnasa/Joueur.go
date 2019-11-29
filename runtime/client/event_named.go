package client

func WaitForEventNamed() string {
	named := ""
	WaitForEvent("named", &named)

	return named
}
