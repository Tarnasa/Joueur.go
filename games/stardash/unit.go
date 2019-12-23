package stardash

// Unit is a unit in the game. May be a corvette, missleboat, martyr,
// transport, miner.
type Unit interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Acted is whether or not this Unit has performed its action this
	// turn.
	Acted() bool

	// DashX is the x value this unit is dashing to.
	DashX() float64

	// DashY is the y value this unit is dashing to.
	DashY() float64

	// Energy is the remaining health of the unit.
	Energy() int64

	// Genarium is the amount of Genarium ore carried by this unit. (0
	// to job carry capacity - other carried items).
	Genarium() int64

	// IsBusy is tracks wheither or not the ship is dashing or Mining.
	// If true, it cannot do anything else.
	IsBusy() bool

	// Job is the Job this Unit has.
	Job() Job

	// Legendarium is the amount of Legendarium ore carried by this
	// unit. (0 to job carry capacity - other carried items).
	Legendarium() int64

	// Moves is the distance this unit can still move.
	Moves() float64

	// Mythicite is the amount of Mythicite carried by this unit. (0
	// to job carry capacity - other carried items).
	Mythicite() int64

	// Owner is the Player that owns and can control this Unit.
	//
	// Value can be returned as a nil pointer.
	Owner() Player

	// Protector is the martyr ship that is currently shielding this
	// ship if any.
	//
	// Value can be returned as a nil pointer.
	Protector() Unit

	// Rarium is the amount of Rarium carried by this unit. (0 to job
	// carry capacity - other carried items).
	Rarium() int64

	// Shield is the sheild that a martyr ship has.
	Shield() int64

	// X is the x value this unit is on.
	X() float64

	// Y is the y value this unit is on.
	Y() float64

	// -- Methods -- \\

	// Attack attacks the specified unit.
	Attack(Unit) bool

	// Dash causes the unit to dash towards the designated
	// destination.
	Dash(float64, float64) bool

	// Mine allows a miner to mine a asteroid
	Mine(Body) bool

	// Move moves this Unit from its current location to the new
	// location specified.
	Move(float64, float64) bool

	// Safe tells you if your ship can move to that location from were
	// it is without clipping the sun.
	Safe(float64, float64) bool

	// Shootdown attacks the specified projectile.
	Shootdown(Projectile) bool

	// Transfer grab materials from a friendly unit. Doesn't use a
	// action.
	Transfer(Unit, int64, string) bool
}
