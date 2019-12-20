package stumped

// Spawner is a resource spawner that generates branches or food.
type Spawner interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// HasBeenHarvested is true if this Spawner has been harvested
	// this turn, and it will not heal at the end of the turn, false
	// otherwise.
	HasBeenHarvested() bool

	// Health is how much health this Spawner has, which is used to
	// calculate how much of its resource can be harvested.
	Health() int64

	// Tile is the Tile this Spawner is on.
	Tile() Tile

	// Type is what type of resource this is ('food' or 'branches').
//
	// Literal Values: "food" or "branches"
	Type() string
}
