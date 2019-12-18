package stumped

// Player is a player in this game. Every AI controls one player.
type Player interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Beavers is the array of Beavers owned by this Player.
	Beavers() []Beaver

	// BranchesToBuildLodge is how many branches are required to build
	// a lodge for this Player.
	BranchesToBuildLodge() int64

	// ClientType is what type of client this is, e.g. 'Python',
	// 'JavaScript', or some other language. For potential data mining
	// purposes.
	ClientType() string

	// Lodges is an array of Tiles that contain lodges owned by this
	// player.
	Lodges() []Tile

	// Lost is if the player lost the game or not.
	Lost() bool

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

	// Won is if the player won the game or not.
	Won() bool
}
