package coreminer

// Bomb is a Bomb in the game.
type Bomb interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Tile is the Tile this Miner is on.
	//
	// Value can be returned as a nil pointer.
	Tile() Tile

	// Timer is the number of turns before this Bomb explodes. Zero
	// means it will explode after the current turn.
	Timer() int64
}
