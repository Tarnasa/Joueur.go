package newtonian

// Player is a player in this game. Every AI controls one player.
type Player interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// ClientType is what type of client this is, e.g. 'Python',
	// 'JavaScript', or some other language. For potential data mining
	// purposes.
	ClientType() string

	// GeneratorTiles is every generator Tile owned by this Player.
	// (arrayed from the outer edges inward, from top to bottom).
	GeneratorTiles() []Tile

	// Heat is the amount of heat this Player has.
	Heat() int64

	// InternSpawn is the time left till a intern spawns. (0 to
	// spawnTime).
	InternSpawn() int64

	// Lost is if the player lost the game or not.
	Lost() bool

	// ManagerSpawn is the time left till a manager spawns. (0 to
	// spawnTime).
	ManagerSpawn() int64

	// Name is the name of the player.
	Name() string

	// Opponent is this player's opponent in the game.
	Opponent() Player

	// PhysicistSpawn is the time left till a physicist spawns. (0 to
	// spawnTime).
	PhysicistSpawn() int64

	// Pressure is the amount of pressure this Player has.
	Pressure() int64

	// ReasonLost is the reason why the player lost the game.
	ReasonLost() string

	// ReasonWon is the reason why the player won the game.
	ReasonWon() string

	// SpawnTiles is all the tiles this Player's units can spawn on.
	// (arrayed from the outer edges inward, from top to bottom).
	SpawnTiles() []Tile

	// TimeRemaining is the amount of time (in ns) remaining for this
	// AI to send commands.
	TimeRemaining() float64

	// Units is every Unit owned by this Player.
	Units() []Unit

	// Won is if the player won the game or not.
	Won() bool
}
