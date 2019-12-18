package necrowar

import "joueur/base"

// Game is send hordes of the undead at your opponent while defending
// yourself against theirs to win.
type Game interface {
	// Parent interfaces
	base.Game

	// -- Attributes -- \\

	// CurrentPlayer is the player whose turn it is currently. That
	// player can send commands. Other players cannot.
	CurrentPlayer() Player

	// CurrentTurn is the current turn number, starting at 0 for the
	// first player's turn.
	CurrentTurn() int64

	// GameObjects is a mapping of every game object's ID to the
	// actual game object. Primarily used by the server and client to
	// easily refer to the game objects via ID.
	GameObjects() map[string]GameObject

	// GoldIncomePerUnit is the amount of gold income per turn per
	// unit in a mine.
	GoldIncomePerUnit() int64

	// IslandIncomePerUnit is the amount of gold income per turn per
	// unit in the island mine.
	IslandIncomePerUnit() int64

	// ManaIncomePerUnit is the Amount of gold income per turn per
	// unit fishing on the river side.
	ManaIncomePerUnit() int64

	// MapHeight is the number of Tiles in the map along the y
	// (vertical) axis.
	MapHeight() int64

	// MapWidth is the number of Tiles in the map along the x
	// (horizontal) axis.
	MapWidth() int64

	// MaxTurns is the maximum number of turns before the game will
	// automatically end.
	MaxTurns() int64

	// Players is array of all the players in the game.
	Players() []Player

	// RiverPhase is the amount of turns it takes between the river
	// changing phases.
	RiverPhase() int64

	// Session is a unique identifier for the game instance that is
	// being played.
	Session() string

	// Tiles is all the tiles in the map, stored in Row-major order.
	// Use `x + y * mapWidth` to access the correct index.
	Tiles() []Tile

	// TimeAddedPerTurn is the amount of time (in nano-seconds) added
	// after each player performs a turn.
	TimeAddedPerTurn() int64

	// TowerJobs is an array of every tower type / job.
	TowerJobs() []TowerJob

	// Towers is every Tower in the game.
	Towers() []Tower

	// UnitJobs is an array of every unit type / job.
	UnitJobs() []UnitJob

	// Units is every Unit in the game.
	Units() []Unit

	// -- Tiled Game Utils -- \\

	// GetTileAt returns the Tile at a give position (x, y).
	GetTileAt(int64, int64) Tile
}
