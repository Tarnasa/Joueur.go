package saloon

import "joueur/base"

// Game is use cowboys to have a good time and play some music on a Piano,
// while brawling with enemy Cowboys.
type Game interface {
	// Parent interfaces
	base.Game

	// -- Attributes -- \\

	// BartenderCooldown is how many turns a Bartender will be busy
	// for after throwing a Bottle.
	BartenderCooldown() int64

	// Bottles is all the beer Bottles currently flying across the
	// saloon in the game.
	Bottles() []Bottle

	// BrawlerDamage is how much damage is applied to neighboring
	// things bit by the Sharpshooter between turns.
	BrawlerDamage() int64

	// Cowboys is every Cowboy in the game.
	Cowboys() []Cowboy

	// CurrentPlayer is the player whose turn it is currently. That
	// player can send commands. Other players cannot.
	CurrentPlayer() Player

	// CurrentTurn is the current turn number, starting at 0 for the
	// first player's turn.
	CurrentTurn() int64

	// Furnishings is every furnishing in the game.
	Furnishings() []Furnishing

	// GameObjects is a mapping of every game object's ID to the
	// actual game object. Primarily used by the server and client to
	// easily refer to the game objects via ID.
	GameObjects() map[string]GameObject

	// Jobs is all the jobs that Cowboys can be called in with.
	Jobs() []string

	// MapHeight is the number of Tiles in the map along the y
	// (vertical) axis.
	MapHeight() int64

	// MapWidth is the number of Tiles in the map along the x
	// (horizontal) axis.
	MapWidth() int64

	// MaxCowboysPerJob is the maximum number of Cowboys a Player can
	// bring into the saloon of each specific job.
	MaxCowboysPerJob() int64

	// MaxTurns is the maximum number of turns before the game will
	// automatically end.
	MaxTurns() int64

	// Players is array of all the players in the game.
	Players() []Player

	// RowdinessToSiesta is when a player's rowdiness reaches or
	// exceeds this number their Cowboys take a collective siesta.
	RowdinessToSiesta() int64

	// Session is a unique identifier for the game instance that is
	// being played.
	Session() string

	// SharpshooterDamage is how much damage is applied to things hit
	// by Sharpshooters when they act.
	SharpshooterDamage() int64

	// SiestaLength is how long siestas are for a player's team.
	SiestaLength() int64

	// Tiles is all the tiles in the map, stored in Row-major order.
	// Use `x + y * mapWidth` to access the correct index.
	Tiles() []Tile

	// TimeAddedPerTurn is the amount of time (in nano-seconds) added
	// after each player performs a turn.
	TimeAddedPerTurn() int64

	// TurnsDrunk is how many turns a Cowboy will be drunk for if a
	// bottle breaks on it.
	TurnsDrunk() int64

	// -- Tiled Game Utils -- \\

	// GetTileAt returns the Tile at a give position (x, y).
	GetTileAt(int64, int64) Tile
}
