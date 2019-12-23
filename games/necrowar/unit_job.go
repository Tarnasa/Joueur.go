package necrowar

// UnitJob is information about a unit's job/type.
type UnitJob interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Damage is the amount of damage this type does per attack.
	Damage() int64

	// GoldCost is how much does this type cost in gold.
	GoldCost() int64

	// Health is the amount of starting health this type has.
	Health() int64

	// ManaCost is how much does this type cost in mana.
	ManaCost() int64

	// Moves is the number of moves this type can make per turn.
	Moves() int64

	// PerTile is how many of this type of unit can take up one tile.
	PerTile() int64

	// Range is amount of tiles away this type has to be in order to
	// be effective.
	Range() int64

	// Title is the type title. 'worker', 'zombie', 'ghoul', 'hound',
	// 'abomination', 'wraith' or 'horseman'.
	//
	// Literal Values: "worker", "zombie", "ghoul", "hound",
	// "abomination", "wraith", or "horseman"
	Title() string
}
