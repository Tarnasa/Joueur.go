// Package errorhandler implements handlers to gracefully capture and log
// errors during game play.
package errorhandler

import (
	"os"
	"runtime/debug"

	"github.com/fatih/color"
)

var errorCodeToNames = map[int]string{
	0:  "NONE",
	20: "INVALID_ARGS",
	21: "COULD_NOT_CONNECT",
	22: "DISCONNECTED_UNEXPECTEDLY",
	23: "CANNOT_READ_SOCKET",
	24: "DELTA_MERGE_FAILURE",
	25: "REFLECTION_FAILED",
	26: "UNKNOWN_EVENT_FROM_SERVER",
	27: "SERVER_TIMEOUT",
	28: "FATAL_EVENT",
	29: "GAME_NOT_FOUND",
	30: "MALFORMED_JSON",
	31: "UNAUTHENTICATED",
	42: "AI_ERRORED",
}

// printErr prints a string as an error string in red.
func printErr(str string) {
	os.Stderr.WriteString(color.RedString(str + "\n"))
}

// ErrorHandler is a callback function that is invoked if HandleError is
// invoked, just before exiting the process. To cleanup stuff like
// disconnecting the client.
var ErrorHandler func()

var handlingErrors = true

// StopHandlingErrors stops the HandleError function from actually handling
// errors and exiting the process.
func StopHandlingErrors() {
	handlingErrors = false
}

// HandleError gracefully handles an unexpected error by printing as much
// information about the error as possible before exiting with a none 0 status
// code.
func HandleError(errorCode int, err error, messages ...string) error {
	if !handlingErrors {
		return err
	}

	if errorCodeName, ok := errorCodeToNames[errorCode]; ok {
		printErr("---\nError: " + errorCodeName)
	}

	for _, message := range messages {
		printErr("---\n" + message)
	}

	if err != nil {
		printErr("---\n" + err.Error())
	}

	stack := debug.Stack()
	if len(stack) > 0 {
		printErr("---\n" + string(stack))
	}

	printErr("---")

	if ErrorHandler != nil {
		ErrorHandler()
	}

	os.Exit(errorCode)

	return err // will never happen, makes compiler happy
}
