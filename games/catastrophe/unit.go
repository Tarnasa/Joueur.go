package catastrophe

// Unit is a unit in the game.
type Unit interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Acted is whether this Unit has performed its action this turn.
	Acted() bool

	// Energy is the amount of energy this Unit has (from 0.0 to
	// 100.0).
	Energy() float64

	// Food is the amount of food this Unit is holding.
	Food() int64

	// Job is the Job this Unit was recruited to do.
	Job() Job

	// Materials is the amount of materials this Unit is holding.
	Materials() int64

	// MovementTarget is the tile this Unit is moving to. This only
	// applies to neutral fresh humans spawned on the road. Otherwise,
	// the tile this Unit is on.
	//
	// Value can be returned as a nil pointer.
	MovementTarget() Tile

	// Moves is how many moves this Unit has left this turn.
	Moves() int64

	// Owner is the Player that owns and can control this Unit, or nil
	// if the Unit is neutral.
	//
	// Value can be returned as a nil pointer.
	Owner() Player

	// Squad is the Units in the same squad as this Unit. Units in the
	// same squad attack and defend together.
	Squad() []Unit

	// Starving is whether this Unit is starving. Starving Units
	// regenerate energy at half the rate they normally would while
	// resting.
	Starving() bool

	// Tile is the Tile this Unit is on.
	//
	// Value can be returned as a nil pointer.
	Tile() Tile

	// TurnsToDie is the number of turns before this Unit dies. This
	// only applies to neutral fresh humans created from combat.
	// Otherwise, 0.
	TurnsToDie() int64

	// -- Methods -- \\

	// Attack attacks an adjacent Tile. Costs an action for each Unit
	// in this Unit's squad. Units in the squad without an action
	// don't participate in combat. Units in combat cannot move
	// afterwards. Attacking structures will not give materials.
	Attack(Tile) bool

	// ChangeJob changes this Unit's Job. Must be at max energy
	// (100.0) to change Jobs.
	ChangeJob(string) bool

	// Construct constructs a Structure on an adjacent Tile.
	Construct(Tile, string) bool

	// Convert converts an adjacent Unit to your side.
	Convert(Tile) bool

	// Deconstruct removes materials from an adjacent Tile's
	// Structure. You cannot deconstruct friendly structures (see
	// Unit.attack).
	Deconstruct(Tile) bool

	// Drop drops some of the given resource on or adjacent to the
	// Unit's Tile. Does not count as an action.
	Drop(Tile, string, int64) bool

	// Harvest harvests the food on an adjacent Tile.
	Harvest(Tile) bool

	// Move moves this Unit from its current Tile to an adjacent Tile.
	Move(Tile) bool

	// Pickup picks up some materials or food on or adjacent to the
	// Unit's Tile. Does not count as an action.
	Pickup(Tile, string, int64) bool

	// Rest regenerates energy. Must be in range of a friendly shelter
	// to rest. Costs an action. Units cannot move after resting.
	Rest() bool
}
