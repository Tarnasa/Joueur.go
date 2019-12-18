package anarchy

// Building is a basic building. It does nothing besides burn down. Other
// Buildings inherit from this class.
type Building interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Bribed is when true this building has already been bribed this
	// turn and cannot be bribed again this turn.
	Bribed() bool

	// BuildingEast is the Building directly to the east of this
	// building, or nil if not present.
	BuildingEast() Building

	// BuildingNorth is the Building directly to the north of this
	// building, or nil if not present.
	BuildingNorth() Building

	// BuildingSouth is the Building directly to the south of this
	// building, or nil if not present.
	BuildingSouth() Building

	// BuildingWest is the Building directly to the west of this
	// building, or nil if not present.
	BuildingWest() Building

	// Fire is how much fire is currently burning the building, and
	// thus how much damage it will take at the end of its owner's
	// turn. 0 means no fire.
	Fire() int64

	// Health is how much health this building currently has. When
	// this reaches 0 the Building has been burned down.
	Health() int64

	// IsHeadquarters is true if this is the Headquarters of the
	// owning player, false otherwise. Burning this down wins the game
	// for the other Player.
	IsHeadquarters() bool

	// Owner is the player that owns this building. If it burns down
	// (health reaches 0) that player gets an additional bribe(s).
	Owner() Player

	// X is the location of the Building along the x-axis.
	X() int64

	// Y is the location of the Building along the y-axis.
	Y() int64
}
