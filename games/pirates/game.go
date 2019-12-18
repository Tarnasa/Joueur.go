package pirates

import "joueur/base"

// Game is steal from merchants and become the most infamous pirate.
type Game interface {
	// Parent interfaces
	base.Game

	// -- Attributes -- \\

	// BuryInterestRate is the rate buried gold increases each turn.
	BuryInterestRate() float64

	// CrewCost is how much gold it costs to construct a single crew.
	CrewCost() int64

	// CrewDamage is how much damage crew deal to each other.
	CrewDamage() int64

	// CrewHealth is the maximum amount of health a crew member can
	// have.
	CrewHealth() int64

	// CrewMoves is the number of moves Units with only crew are given
	// each turn.
	CrewMoves() int64

	// CrewRange is a crew's attack range. Range is circular.
	CrewRange() float64

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

	// HealFactor is how much health a Unit recovers when they rest.
	HealFactor() float64

	// MapHeight is the number of Tiles in the map along the y
	// (vertical) axis.
	MapHeight() int64

	// MapWidth is the number of Tiles in the map along the x
	// (horizontal) axis.
	MapWidth() int64

	// MaxTurns is the maximum number of turns before the game will
	// automatically end.
	MaxTurns() int64

	// MerchantGoldRate is how much gold merchant Ports get each turn.
	MerchantGoldRate() float64

	// MerchantInterestRate is when a merchant ship spawns, the amount
	// of additional gold it has relative to the Port's investment.
	MerchantInterestRate() float64

	// MinInterestDistance is the Euclidean distance buried gold must
	// be from the Player's Port to accumulate interest.
	MinInterestDistance() float64

	// Players is array of all the players in the game.
	Players() []Player

	// Ports is every Port in the game. Merchant ports have owner set
	// to nil.
	Ports() []Port

	// RestRange is how far a Unit can be from a Port to rest. Range
	// is circular.
	RestRange() float64

	// Session is a unique identifier for the game instance that is
	// being played.
	Session() string

	// ShipCost is how much gold it costs to construct a ship.
	ShipCost() int64

	// ShipDamage is how much damage ships deal to ships and ports.
	ShipDamage() int64

	// ShipHealth is the maximum amount of health a ship can have.
	ShipHealth() int64

	// ShipMoves is the number of moves Units with ships are given
	// each turn.
	ShipMoves() int64

	// ShipRange is a ship's attack range. Range is circular.
	ShipRange() float64

	// Tiles is all the tiles in the map, stored in Row-major order.
	// Use `x + y * mapWidth` to access the correct index.
	Tiles() []Tile

	// TimeAddedPerTurn is the amount of time (in nano-seconds) added
	// after each player performs a turn.
	TimeAddedPerTurn() int64

	// Units is every Unit in the game. Merchant units have targetPort
	// set to a port.
	Units() []Unit

	// -- Tiled Game Utils -- \\

	// GetTileAt returns the Tile at a give position (x, y).
	GetTileAt(int64, int64) Tile
}
