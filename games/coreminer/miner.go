package coreminer

// Miner is a Miner in the game.
type Miner interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Bombs is the number of bombs being carried by this Miner.
	Bombs() int64

	// BuildingMaterials is the number of building materials carried
	// by this Miner.
	BuildingMaterials() int64

	// CurrentUpgrade is the Upgrade this Miner is on.
	CurrentUpgrade() Upgrade

	// Dirt is the amount of dirt carried by this Miner.
	Dirt() int64

	// Health is the remaining health of this Miner.
	Health() int64

	// MiningPower is the remaining mining power this Miner has this
	// turn.
	MiningPower() int64

	// Moves is the number of moves this Miner has left this turn.
	Moves() int64

	// Ore is the amount of ore carried by this Miner.
	Ore() int64

	// Owner is the Player that owns and can control this Miner.
	Owner() Player

	// Tile is the Tile this Miner is on.
	//
	// Value can be returned as a nil pointer.
	Tile() Tile

	// UpgradeLevel is the upgrade level of this Miner. Starts at 0.
	UpgradeLevel() int64

	// -- Methods -- \\

	// Build builds a support, shield, or ladder on Miner's Tile, or
	// an adjacent Tile.
	Build(Tile, string) bool

	// Buy purchase a resource from the Player's base or hopper.
	Buy(string, int64) bool

	// Dump dumps materials from cargo to an adjacent Tile. If the
	// Tile is a base or a hopper Tile, materials are sold instead of
	// placed.
	Dump(Tile, string, int64) bool

	// Mine mines the Tile the Miner is on or an adjacent Tile.
	Mine(Tile, int64) bool

	// Move moves this Miner from its current Tile to an adjacent
	// Tile.
	Move(Tile) bool

	// Transfer transfers a resource from the one Miner to another.
	Transfer(Miner, string, int64) bool

	// Upgrade upgrade this Miner by installing an upgrade module.
	Upgrade() bool
}
