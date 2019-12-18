package necrowar

// Tower is a tower in the game. Used to combat enemy waves.
type Tower interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Attacked is whether this tower has attacked this turn or not.
	Attacked() bool

	// Cooldown is how many turns are left before it can fire again.
	Cooldown() int64

	// Health is how much remaining health this tower has.
	Health() int64

	// Job is what type of tower this is (it's job).
	Job() TowerJob

	// Owner is the player that built / owns this tower.
	Owner() Player

	// Tile is the Tile this Tower is on.
	Tile() Tile

	// -- Methods -- \\

	// Attack attacks an enemy unit on an tile within it's range.
	Attack(Tile) bool
}
