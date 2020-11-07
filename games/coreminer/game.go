package coreminer

import "joueur/base"

// Game is mine resources to obtain more value than your opponent.
type Game interface {
	// Parent interfaces
	base.Game

	// -- Attributes -- \\

	// BombPrice is the monetary price of a bomb when bought or sold.
	BombPrice() int64

	// BombSize is the amount of cargo space taken up by a Bomb.
	BombSize() int64

	// Bombs is every Bomb in the game.
	Bombs() []Bomb

	// BuildingMaterialPrice is the monetary price of building
	// materials when bought.
	BuildingMaterialPrice() int64

	// CurrentPlayer is the player whose turn it is currently. That
	// player can send commands. Other players cannot.
	CurrentPlayer() Player

	// CurrentTurn is the current turn number, starting at 0 for the
	// first player's turn.
	CurrentTurn() int64

	// DirtPrice is the monetary price of dirt when bought or sold.
	DirtPrice() int64

	// FallDamage is the amount of damage taken per Tile fallen.
	FallDamage() int64

	// FallWeightDamage is the amount of extra damage taken for
	// falling while carrying a large amount of cargo.
	FallWeightDamage() int64

	// GameObjects is a mapping of every game object's ID to the
	// actual game object. Primarily used by the server and client to
	// easily refer to the game objects via ID.
	GameObjects() map[string]GameObject

	// LadderCost is the amount of building material required to build
	// a ladder.
	LadderCost() int64

	// LadderHealth is the amount of mining power needed to remove a
	// ladder from a Tile.
	LadderHealth() int64

	// LargeCargoSize is the amount deemed as a large amount of cargo.
	LargeCargoSize() int64

	// LargeMaterialSize is the amount deemed as a large amount of
	// material.
	LargeMaterialSize() int64

	// MapHeight is the number of Tiles in the map along the y
	// (vertical) axis.
	MapHeight() int64

	// MapWidth is the number of Tiles in the map along the x
	// (horizontal) axis.
	MapWidth() int64

	// MaxShielding is the maximum amount of shielding possible on a
	// Tile.
	MaxShielding() int64

	// MaxTurns is the maximum number of turns before the game will
	// automatically end.
	MaxTurns() int64

	// MaxUpgradeLevel is the highest upgrade level allowed on a
	// Miner.
	MaxUpgradeLevel() int64

	// Miners is every Miner in the game.
	Miners() []Miner

	// OrePrice is the amount of money awarded when ore is dumped in
	// the base and sold.
	OrePrice() int64

	// OreValue is the amount of value awarded when ore is dumped in
	// the base and sold.
	OreValue() int64

	// Players is array of all the players in the game.
	Players() []Player

	// Session is a unique identifier for the game instance that is
	// being played.
	Session() string

	// ShieldCost is the amount of building material required to
	// shield a Tile.
	ShieldCost() int64

	// ShieldHealth is the amount of mining power needed to remove one
	// unit of shielding off a Tile.
	ShieldHealth() int64

	// SpawnPrice is the monetary price of spawning a Miner.
	SpawnPrice() int64

	// SuffocationDamage is the amount of damage taken when
	// suffocating inside a filled Tile.
	SuffocationDamage() int64

	// SuffocationWeightDamage is the amount of extra damage taken for
	// suffocating under a large amount of material.
	SuffocationWeightDamage() int64

	// SupportCost is the amount of building material required to
	// build a support.
	SupportCost() int64

	// SupportHealth is the amount of mining power needed to remove a
	// support from a Tile.
	SupportHealth() int64

	// Tiles is all the tiles in the map, stored in Row-major order.
	// Use `x + y * mapWidth` to access the correct index.
	Tiles() []Tile

	// TimeAddedPerTurn is the amount of time (in nano-seconds) added
	// after each player performs a turn.
	TimeAddedPerTurn() float64

	// UpgradePrice is the cost to upgrade a Miner.
	UpgradePrice() int64

	// Upgrades is every Upgrade for a Miner in the game.
	Upgrades() []Upgrade

	// VictoryAmount is the amount of victory points (value) required
	// to win.
	VictoryAmount() int64

	// -- Tiled Game Utils -- \\

	// GetTileAt returns the Tile at a give position (x, y).
	GetTileAt(int64, int64) Tile
}
