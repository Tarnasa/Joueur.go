package necrowar

// Unit is a unit in the game. May be a worker, zombie, ghoul, hound,
// abomination, wraith or horseman.
type Unit interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Acted is whether or not this Unit has performed its action this
	// turn (attack or build).
	Acted() bool

	// Health is the remaining health of a unit.
	Health() int64

	// Job is the type of unit this is.
	Job() UnitJob

	// Moves is the number of moves this unit has left this turn.
	Moves() int64

	// Owner is the Player that owns and can control this Unit.
//
	// Value can be returned as a nil pointer.
	Owner() Player

	// Tile is the Tile this Unit is on.
//
	// Value can be returned as a nil pointer.
	Tile() Tile

	// -- Methods -- \\

	// Attack attacks an enemy tower on an adjacent tile.
	Attack(Tile) bool

	// Build unit, if it is a worker, builds a tower on the tile it is
	// on, only workers can do this.
	Build(string) bool

	// Fish stops adjacent to a river tile and begins fishing for
	// mana.
	Fish(Tile) bool

	// Mine enters a mine and is put to work gathering resources.
	Mine(Tile) bool

	// Move moves this Unit from its current Tile to an adjacent Tile.
	Move(Tile) bool
}
