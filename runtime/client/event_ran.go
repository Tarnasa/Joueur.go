package client

func WaitForEventRan() interface{} {
	var returned interface{}
	WaitForEvent("ran", &returned)

	return returned
}
