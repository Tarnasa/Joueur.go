package stumped

import "joueur/base"

// Game is gather branches and build up your lodge as beavers fight to
// survive.
type Game interface {
	// Parent interfaces
	base.Game

	// -- Attributes -- \\

	// Beavers is every Beaver in the game.
	Beavers() []Beaver

	// CurrentPlayer is the player whose turn it is currently. That
	// player can send commands. Other players cannot.
	CurrentPlayer() Player

	// CurrentTurn is the current turn number, starting at 0 for the
	// first player's turn.
	CurrentTurn() int64

	// FreeBeaversCount is when a Player has less Beavers than this
	// number, then recruiting other Beavers is free.
	FreeBeaversCount() int64

	// GameObjects is a mapping of every game object's ID to the
	// actual game object. Primarily used by the server and client to
	// easily refer to the game objects via ID.
	GameObjects() map[string]GameObject

	// Jobs is all the Jobs that Beavers can have in the game.
	Jobs() []Job

	// LodgeCostConstant is constant number used to calculate what it
	// costs to spawn a new lodge.
	LodgeCostConstant() float64

	// LodgesToWin is how many lodges must be owned by a Player at
	// once to win the game.
	LodgesToWin() int64

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

	// Session is a unique identifier for the game instance that is
	// being played.
	Session() string

	// Spawner is every Spawner in the game.
	Spawner() []Spawner

	// SpawnerHarvestConstant is constant number used to calculate how
	// many branches/food Beavers harvest from Spawners.
	SpawnerHarvestConstant() float64

	// SpawnerTypes is all the types of Spawners in the game.
	SpawnerTypes() []string

	// Tiles is all the tiles in the map, stored in Row-major order.
	// Use `x + y * mapWidth` to access the correct index.
	Tiles() []Tile

	// TimeAddedPerTurn is the amount of time (in nano-seconds) added
	// after each player performs a turn.
	TimeAddedPerTurn() int64

	// -- Tiled Game Utils -- \\

	// GetTileAt returns the Tile at a give position (x, y).
	GetTileAt(int64, int64) Tile
}
