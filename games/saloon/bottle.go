package saloon

// Bottle is a bottle thrown by a bartender at a Tile.
type Bottle interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Direction is the Direction this Bottle is flying and will move
	// to between turns, can be 'North', 'East', 'South', or 'West'.
//
	// Literal Values: "North", "East", "South", or "West"
	Direction() string

	// DrunkDirection is the direction any Cowboys hit by this will
	// move, can be 'North', 'East', 'South', or 'West'.
//
	// Literal Values: "North", "East", "South", or "West"
	DrunkDirection() string

	// IsDestroyed is true if this Bottle has impacted and has been
	// destroyed (removed from the Game). False if still in the game
	// flying through the saloon.
	IsDestroyed() bool

	// Tile is the Tile this bottle is currently flying over.
//
	// Value can be returned as a nil pointer.
	Tile() Tile
}
