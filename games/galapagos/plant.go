package galapagos

// Plant is a Plant in the game.
type Plant interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// GrowthRate is the total number of turns it takes this plant to
	// grow in size.
	GrowthRate() int64

	// Size is the size of the plant.
	Size() int64

	// Tile is the Tile this Plant occupies.
	//
	// Value can be returned as a nil pointer.
	Tile() Tile

	// TurnsUntilGrowth is the number of turns left until this plant
	// will grow again.
	TurnsUntilGrowth() int64
}
