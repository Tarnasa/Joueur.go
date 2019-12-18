package necrowar

// Player is a player in this game. Every AI controls one player.
type Player interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// ClientType is what type of client this is, e.g. 'Python',
	// 'JavaScript', or some other language. For potential data mining
	// purposes.
	ClientType() string

	// Gold is the amount of gold this Player has.
	Gold() int64

	// Health is the amount of health remaining for this player's main
	// unit.
	Health() int64

	// HomeBase is the tile that the home base is located on.
	HomeBase() []Tile

	// Lost is if the player lost the game or not.
	Lost() bool

	// Mana is the amount of mana this player has.
	Mana() int64

	// Name is the name of the player.
	Name() string

	// Opponent is this player's opponent in the game.
	Opponent() Player

	// ReasonLost is the reason why the player lost the game.
	ReasonLost() string

	// ReasonWon is the reason why the player won the game.
	ReasonWon() string

	// Side is all tiles that this player can build on and move
	// workers on.
	Side() []Tile

	// TimeRemaining is the amount of time (in ns) remaining for this
	// AI to send commands.
	TimeRemaining() float64

	// Towers is every Tower owned by this player.
	Towers() []Tower

	// Units is every Unit owned by this Player.
	Units() []Unit

	// Won is if the player won the game or not.
	Won() bool
}
