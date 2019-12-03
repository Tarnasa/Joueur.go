package chess

// A player in this game. Every AI controls one player.
type Player interface {
	// Parent interfaces
	GameObject
	// -- Attributes -- \\

	// What type of client this is, e.g. 'Python', 'JavaScript', or some other language. For potential data mining purposes.
	ClientType() string

	// The color (side) of this player. Either 'white' or 'black', with the 'white' player having the first move.
	Color() string

	// If the player lost the game or not.
	Lost() bool

	// The name of the player.
	Name() string

	// This player's opponent in the game.
	Opponent() Player

	// The reason why the player lost the game.
	ReasonLost() string

	// The reason why the player won the game.
	ReasonWon() string

	// The amount of time (in ns) remaining for this AI to send commands.
	TimeRemaining() float64

	// If the player won the game or not.
	Won() bool

}
