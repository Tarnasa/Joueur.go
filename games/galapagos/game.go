package galapagos

import "joueur/base"

// Game is adapt, Evolve, Segfault.
type Game interface {
	// Parent interfaces
	base.Game

	// -- Attributes -- \\

	// BaseHealth is the amount of health that a creature with a 0
	// endurance stat starts with.
	BaseHealth() int64

	// Creatures is every Creature in the game.
	Creatures() []Creature

	// CurrentPlayer is the player whose turn it is currently. That
	// player can send commands. Other players cannot.
	CurrentPlayer() Player

	// CurrentTurn is the current turn number, starting at 0 for the
	// first player's turn.
	CurrentTurn() int64

	// DamageMultiplier is how much to damage an opponent per
	// difference of carnivorism and defense.
	DamageMultiplier() int64

	// GameObjects is a mapping of every game object's ID to the
	// actual game object. Primarily used by the server and client to
	// easily refer to the game objects via ID.
	GameObjects() map[string]GameObject

	// HealthPerBreed is the amount of extra health from both breeding
	// creatures required if you have more total health than your
	// opponent.
	HealthPerBreed() int64

	// HealthPerCarnivorism is multiplied by carnivorism to determine
	// health gained from eating creatures.
	HealthPerCarnivorism() int64

	// HealthPerEndurance is the amount of extra health for each point
	// of endurance.
	HealthPerEndurance() int64

	// HealthPerHerbivorism is multiplied by herbivorism to determine
	// health gained from biting plants.
	HealthPerHerbivorism() int64

	// HealthPerMove is the amount of health required to move.
	HealthPerMove() int64

	// HealthPerTurn is the amount of health lost after each of your
	// turns.
	HealthPerTurn() int64

	// MapHeight is the number of Tiles in the map along the y
	// (vertical) axis.
	MapHeight() int64

	// MapWidth is the number of Tiles in the map along the x
	// (horizontal) axis.
	MapWidth() int64

	// MaxPlantSize is the maximum size a plant to grow to.
	MaxPlantSize() int64

	// MaxStartingCreatures is the maximum number of creatures that
	// each player will start with.
	MaxStartingCreatures() int64

	// MaxStartingPlants is the maximum number of plants that the map
	// will start with.
	MaxStartingPlants() int64

	// MaxStatValue is the maxmimum value that a stat (carnivorism,
	// herbivorism, defense, endurance, speed) can have.
	MaxStatValue() int64

	// MaxTurns is the maximum number of turns before the game will
	// automatically end.
	MaxTurns() int64

	// MinStartingCreatures is the minimum number of creatures that
	// each player will start with.
	MinStartingCreatures() int64

	// MinStartingPlants is the minimum number of plants that the map
	// will start with.
	MinStartingPlants() int64

	// Plants is every Plant in the game.
	Plants() []Plant

	// Players is array of all the players in the game.
	Players() []Player

	// Session is a unique identifier for the game instance that is
	// being played.
	Session() string

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
