package coreminer

// Player is a player in this game. Every AI controls one player.
type Player interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// BaseTile is the Tile this Player's base is on.
	BaseTile() Tile

	// Bombs is every Bomb owned by this Player.
	Bombs() []Bomb

	// ClientType is what type of client this is, e.g. 'Python',
	// 'JavaScript', or some other language. For potential data mining
	// purposes.
	ClientType() string

	// HopperTiles is the Tiles this Player's hoppers are on.
	HopperTiles() []Tile

	// Lost is if the player lost the game or not.
	Lost() bool

	// Miners is every Miner owned by this Player.
	Miners() []Miner

	// Money is the amount of money this Player currently has.
	Money() int64

	// Name is the name of the player.
	Name() string

	// Opponent is this player's opponent in the game.
	Opponent() Player

	// ReasonLost is the reason why the player lost the game.
	ReasonLost() string

	// ReasonWon is the reason why the player won the game.
	ReasonWon() string

	// TimeRemaining is the amount of time (in ns) remaining for this
	// AI to send commands.
	TimeRemaining() float64

	// Value is the amount of value (victory points) this Player has
	// gained.
	Value() int64

	// Won is if the player won the game or not.
	Won() bool

	// -- Methods -- \\

	// SpawnMiner spawns a Miner on this Player's base Tile.
	SpawnMiner() bool
}
