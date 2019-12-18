package checkers

import "joueur/base"

// Game is the simple version of American Checkers. An 8x8 board with 12
// checkers on each side that must move diagonally to the opposing side
// until kinged.
type Game interface {
	// Parent interfaces
	base.Game

	// -- Attributes -- \\

	// BoardHeight is the height of the board for the Y component of a
	// checker.
	BoardHeight() int64

	// BoardWidth is the width of the board for X component of a
	// checker.
	BoardWidth() int64

	// CheckerMoved is the checker that last moved and must be moved
	// because only one checker can move during each players turn.
	CheckerMoved() Checker

	// CheckerMovedJumped is if the last checker that moved jumped,
	// meaning it can move again.
	CheckerMovedJumped() bool

	// Checkers is all the checkers currently in the game.
	Checkers() []Checker

	// CurrentPlayer is the player whose turn it is currently. That
	// player can send commands. Other players cannot.
	CurrentPlayer() Player

	// CurrentTurn is the current turn number, starting at 0 for the
	// first player's turn.
	CurrentTurn() int64

	// GameObjects is a mapping of every game object's ID to the
	// actual game object. Primarily used by the server and client to
	// easily refer to the game objects via ID.
	GameObjects() map[string]GameObject

	// MaxTurns is the maximum number of turns before the game will
	// automatically end.
	MaxTurns() int64

	// Players is array of all the players in the game.
	Players() []Player

	// Session is a unique identifier for the game instance that is
	// being played.
	Session() string

	// TimeAddedPerTurn is the amount of time (in nano-seconds) added
	// after each player performs a turn.
	TimeAddedPerTurn() int64
}
