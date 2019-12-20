package newtonian

// Unit is a unit in the game. May be a manager, intern, or physicist.
type Unit interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Acted is whether or not this Unit has performed its action this
	// turn.
	Acted() bool

	// Blueium is the amount of blueium carried by this unit. (0 to
	// job carry capacity - other carried items).
	Blueium() int64

	// BlueiumOre is the amount of blueium ore carried by this unit.
	// (0 to job carry capacity - other carried items).
	BlueiumOre() int64

	// Health is the remaining health of a unit.
	Health() int64

	// Job is the Job this Unit has.
	Job() Job

	// Moves is the number of moves this unit has left this turn.
	Moves() int64

	// Owner is the Player that owns and can control this Unit.
//
	// Value can be returned as a nil pointer.
	Owner() Player

	// Redium is the amount of redium carried by this unit. (0 to job
	// carry capacity - other carried items).
	Redium() int64

	// RediumOre is the amount of redium ore carried by this unit. (0
	// to job carry capacity - other carried items).
	RediumOre() int64

	// StunImmune is duration of stun immunity. (0 to timeImmune).
	StunImmune() int64

	// StunTime is duration the unit is stunned. (0 to the game
	// constant stunTime).
	StunTime() int64

	// Tile is the Tile this Unit is on.
//
	// Value can be returned as a nil pointer.
	Tile() Tile

	// -- Methods -- \\

	// Act makes the unit do something to a machine or unit adjacent
	// to its tile. Interns sabotage, physicists work. Interns stun
	// physicist, physicist stuns manager, manager stuns intern.
	Act(Tile) bool

	// Attack attacks a unit on an adjacent tile.
	Attack(Tile) bool

	// Drop drops materials at the units feet or adjacent tile.
	Drop(Tile, int64, string) bool

	// Move moves this Unit from its current Tile to an adjacent Tile.
	Move(Tile) bool

	// Pickup picks up material at the units feet or adjacent tile.
	Pickup(Tile, int64, string) bool
}
