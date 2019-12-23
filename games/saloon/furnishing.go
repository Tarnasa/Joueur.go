package saloon

// Furnishing is an furnishing in the Saloon that must be pathed around, or
// destroyed.
type Furnishing interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Health is how much health this Furnishing currently has.
	Health() int64

	// IsDestroyed is if this Furnishing has been destroyed, and has
	// been removed from the game.
	IsDestroyed() bool

	// IsPiano is true if this Furnishing is a piano and can be
	// played, False otherwise.
	IsPiano() bool

	// IsPlaying is if this is a piano and a Cowboy is playing it this
	// turn.
	IsPlaying() bool

	// Tile is the Tile that this Furnishing is located on.
	//
	// Value can be returned as a nil pointer.
	Tile() Tile
}
