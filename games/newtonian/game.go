package newtonian

import "joueur/base"

// Game is combine elements and be the first scientists to create fusion.
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

	// InternCap is the maximum number of interns a player can have.
	InternCap() int64

	// Jobs is an array of all jobs. The first element is intern,
	// second is physicists, and third is manager.
	Jobs() []Job

	// Machines is every Machine in the game.
	Machines() []Machine

	// ManagerCap is the maximum number of managers a player can have.
	ManagerCap() int64

	// MapHeight is the number of Tiles in the map along the y
	// (vertical) axis.
	MapHeight() int64

	// MapWidth is the number of Tiles in the map along the x
	// (horizontal) axis.
	MapWidth() int64

	// MaterialSpawn is the number of materials that spawn per spawn
	// cycle.
	MaterialSpawn() int64

	// MaxTurns is the maximum number of turns before the game will
	// automatically end.
	MaxTurns() int64

	// PhysicistCap is the maximum number of physicists a player can
	// have.
	PhysicistCap() int64

	// Players is array of all the players in the game.
	Players() []Player

	// RefinedValue is the amount of victory points added when a
	// refined ore is consumed by the generator.
	RefinedValue() int64

	// RegenerateRate is the percent of max HP regained when a unit
	// end their turn on a tile owned by their player.
	RegenerateRate() float64

	// Session is a unique identifier for the game instance that is
	// being played.
	Session() string

	// SpawnTime is the amount of turns it takes a unit to spawn.
	SpawnTime() int64

	// StunTime is the amount of turns a unit cannot do anything when
	// stunned.
	StunTime() int64

	// Tiles is all the tiles in the map, stored in Row-major order.
	// Use `x + y * mapWidth` to access the correct index.
	Tiles() []Tile

	// TimeAddedPerTurn is the amount of time (in nano-seconds) added
	// after each player performs a turn.
	TimeAddedPerTurn() int64

	// TimeImmune is the number turns a unit is immune to being
	// stunned.
	TimeImmune() int64

	// Units is every Unit in the game.
	Units() []Unit

	// VictoryAmount is the amount of combined heat and pressure that
	// you need to win.
	VictoryAmount() int64

	// -- Tiled Game Utils -- \\

	// GetTileAt returns the Tile at a give position (x, y).
	GetTileAt(int64, int64) Tile
}
