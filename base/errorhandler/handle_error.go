package errorhandler

import (
	"os"

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

func printErr(str string, a ...interface{}) {
	os.Stderr.WriteString(color.RedString(str+"\n", a...))
}

var errorHandler = func() {}

func RegisterErrorHandler(handler func()) {
	errorHandler = handler
}

func HandleError(errorCode int, err error, messages ...string) {
	if errorCodeName, ok := errorCodeToNames[errorCode]; ok {
		printErr("---\nError: "+ errorCodeName)
	}

	for _, message := range messages {
		printErr("---\n" + message)
	}

	if err != nil {
		printErr("---\n" + err.Error())
	}

	printErr("---")

	if errorHandler != nil {
		errorHandler()
	}

	os.Exit(errorCode)
}
