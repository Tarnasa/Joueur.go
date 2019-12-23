package stumped

// Beaver is a beaver in the game.
type Beaver interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Actions is the number of actions remaining for the Beaver this
	// turn.
	Actions() int64

	// Branches is the amount of branches this Beaver is holding.
	Branches() int64

	// Food is the amount of food this Beaver is holding.
	Food() int64

	// Health is how much health this Beaver has left.
	Health() int64

	// Job is the Job this Beaver was recruited to do.
	Job() Job

	// Moves is how many moves this Beaver has left this turn.
	Moves() int64

	// Owner is the Player that owns and can control this Beaver.
	Owner() Player

	// Recruited is true if the Beaver has finished being recruited
	// and can do things, False otherwise.
	Recruited() bool

	// Tile is the Tile this Beaver is on.
	//
	// Value can be returned as a nil pointer.
	Tile() Tile

	// TurnsDistracted is number of turns this Beaver is distracted
	// for (0 means not distracted).
	TurnsDistracted() int64

	// -- Methods -- \\

	// Attack attacks another adjacent beaver.
	Attack(Beaver) bool

	// BuildLodge builds a lodge on the Beavers current Tile.
	BuildLodge() bool

	// Drop drops some of the given resource on the beaver's Tile.
	Drop(Tile, string, int64) bool

	// Harvest harvests the branches or food from a Spawner on an
	// adjacent Tile.
	Harvest(Spawner) bool

	// Move moves this Beaver from its current Tile to an adjacent
	// Tile.
	Move(Tile) bool

	// Pickup picks up some branches or food on the beaver's tile.
	Pickup(Tile, string, int64) bool
}
