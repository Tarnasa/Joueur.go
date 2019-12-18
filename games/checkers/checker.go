package checkers

// Checker is a checker on the game board.
type Checker interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Kinged is if the checker has been kinged and can move
	// backwards.
	Kinged() bool

	// Owner is the player that controls this Checker.
	Owner() Player

	// X is the x coordinate of the checker.
	X() int64

	// Y is the y coordinate of the checker.
	Y() int64

	// -- Methods -- \\

	// IsMine returns if the checker is owned by your player or not.
	IsMine() bool

	// Move moves the checker from its current location to the given
	// (x, y).
	Move(int64, int64) Checker
}
