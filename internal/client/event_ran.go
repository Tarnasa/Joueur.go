package client

// WaitForEventRan waits for the "ran" event and returns the event's data.
func WaitForEventRan() interface{} {
	var returned interface{}
	waitForEvent("ran", &returned)

	return returned
}
