package saloon

// Cowboy is a person on the map that can move around and interact within
// the saloon.
type Cowboy interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// CanMove is if the Cowboy can be moved this turn via its owner.
	CanMove() bool

	// DrunkDirection is the direction this Cowboy is moving while
	// drunk. Will be 'North', 'East', 'South', or 'West' when drunk;
	// or '' (empty string) when not drunk.
//
	// Literal Values: "", "North", "East", "South", or "West"
	DrunkDirection() string

	// Focus is how much focus this Cowboy has. Different Jobs do
	// different things with their Cowboy's focus.
	Focus() int64

	// Health is how much health this Cowboy currently has.
	Health() int64

	// IsDead is if this Cowboy is dead and has been removed from the
	// game.
	IsDead() bool

	// IsDrunk is if this Cowboy is drunk, and will automatically
	// walk.
	IsDrunk() bool

	// Job is the job that this Cowboy does, and dictates how they
	// fight and interact within the Saloon.
//
	// Literal Values: "Bartender", "Brawler", or "Sharpshooter"
	Job() string

	// Owner is the Player that owns and can control this Cowboy.
	Owner() Player

	// Tile is the Tile that this Cowboy is located on.
//
	// Value can be returned as a nil pointer.
	Tile() Tile

	// Tolerance is how many times this unit has been drunk before
	// taking their siesta and reseting this to 0.
	Tolerance() int64

	// TurnsBusy is how many turns this unit has remaining before it
	// is no longer busy and can `act()` or `play()` again.
	TurnsBusy() int64

	// -- Methods -- \\

	// Act does their job's action on a Tile.
	Act(Tile, string) bool

	// Move moves this Cowboy from its current Tile to an adjacent
	// Tile.
	Move(Tile) bool

	// Play sits down and plays a piano.
	Play(Furnishing) bool
}
