package catastrophe

import "joueur/base"

// Game is convert as many humans to as you can to survive in this post-
// apocalyptic wasteland.
type Game interface {
	// Parent interfaces
	base.Game

	// -- Attributes -- \\

	// CatEnergyMult is the multiplier for the amount of energy
	// regenerated when resting in a shelter with the cat overlord.
	CatEnergyMult() float64

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

	// HarvestCooldown is the amount of turns it takes for a Tile that
	// was just harvested to grow food again.
	HarvestCooldown() int64

	// Jobs is all the Jobs that Units can have in the game.
	Jobs() []Job

	// LowerHarvestAmount is the amount that the harvest rate is
	// lowered each season.
	LowerHarvestAmount() int64

	// MapHeight is the number of Tiles in the map along the y
	// (vertical) axis.
	MapHeight() int64

	// MapWidth is the number of Tiles in the map along the x
	// (horizontal) axis.
	MapWidth() int64

	// MaxTurns is the maximum number of turns before the game will
	// automatically end.
	MaxTurns() int64

	// MonumentCostMult is the multiplier for the cost of actions when
	// performing them in range of a monument. Does not effect pickup
	// cost.
	MonumentCostMult() float64

	// MonumentMaterials is the number of materials in a monument.
	MonumentMaterials() int64

	// NeutralMaterials is the number of materials in a neutral
	// Structure.
	NeutralMaterials() int64

	// Players is array of all the players in the game.
	Players() []Player

	// Session is a unique identifier for the game instance that is
	// being played.
	Session() string

	// ShelterMaterials is the number of materials in a shelter.
	ShelterMaterials() int64

	// StartingFood is the amount of food Players start with.
	StartingFood() int64

	// StarvingEnergyMult is the multiplier for the amount of energy
	// regenerated when resting while starving.
	StarvingEnergyMult() float64

	// Structures is every Structure in the game.
	Structures() []Structure

	// Tiles is all the tiles in the map, stored in Row-major order.
	// Use `x + y * mapWidth` to access the correct index.
	Tiles() []Tile

	// TimeAddedPerTurn is the amount of time (in nano-seconds) added
	// after each player performs a turn.
	TimeAddedPerTurn() int64

	// TurnsBetweenHarvests is after a food tile is harvested, the
	// number of turns before it can be harvested again.
	TurnsBetweenHarvests() int64

	// TurnsToCreateHuman is the number of turns between fresh humans
	// being spawned on the road.
	TurnsToCreateHuman() int64

	// TurnsToLowerHarvest is the number of turns before the harvest
	// rate is lowered (length of each season basically).
	TurnsToLowerHarvest() int64

	// Units is every Unit in the game.
	Units() []Unit

	// WallMaterials is the number of materials in a wall.
	WallMaterials() int64

	// -- Tiled Game Utils -- \\

	// GetTileAt returns the Tile at a give position (x, y).
	GetTileAt(int64, int64) Tile
}
