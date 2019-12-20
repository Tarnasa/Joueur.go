package necrowar

// TowerJob is information about a tower's job/type.
type TowerJob interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// AllUnits is whether this tower type hits all of the units on a
	// tile (true) or one at a time (false).
	AllUnits() bool

	// Damage is the amount of damage this type does per attack.
	Damage() int64

	// GoldCost is how much does this type cost in gold.
	GoldCost() int64

	// Health is the amount of starting health this type has.
	Health() int64

	// ManaCost is how much does this type cost in mana.
	ManaCost() int64

	// Range is the number of tiles this type can attack from.
	Range() int64

	// Title is the type title. 'arrow', 'aoe', 'balarraya',
	// 'cleansing', or 'castle'.
//
	// Literal Values: "arrow", "aoe", "ballista", "cleansing", or
	// "castle"
	Title() string

	// TurnsBetweenAttacks is how many turns have to take place
	// between this type's attacks.
	TurnsBetweenAttacks() int64
}
