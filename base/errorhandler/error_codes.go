package errorhandler

const (
	None                     = 0
	InvalidArgs              = 20
	CouldNotConnect          = 21
	DisconnectedUnexpectedly = 22
	CannotReadSocket         = 23
	DeltaMergeFailure        = 24
	ReflectionFailed         = 25
	UnknownEventFromServer   = 26
	ServerTimeout            = 27
	FatalEvent               = 28
	GameNotFound             = 29
	MalformedJSON            = 30
	Unauthenticated          = 31
	AIErrored                = 42
)
