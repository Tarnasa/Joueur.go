package stardash

// Player is a player in this game. Every AI controls one player.
type Player interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// ClientType is what type of client this is, e.g. 'Python',
	// 'JavaScript', or some other language. For potential data mining
	// purposes.
	ClientType() string

	// HomeBase is the home base of the player.
	HomeBase() Body

	// Lost is if the player lost the game or not.
	Lost() bool

	// Money is the amount of money this Player has.
	Money() int64

	// Name is the name of the player.
	Name() string

	// Opponent is this player's opponent in the game.
	Opponent() Player

	// Projectiles is every Projectile owned by this Player. The
	// earlier in the array the older they are.
	Projectiles() []Projectile

	// ReasonLost is the reason why the player lost the game.
	ReasonLost() string

	// ReasonWon is the reason why the player won the game.
	ReasonWon() string

	// TimeRemaining is the amount of time (in ns) remaining for this
	// AI to send commands.
	TimeRemaining() float64

	// Units is every Unit owned by this Player. The earlier in the
	// array the older they are.
	Units() []Unit

	// VictoryPoints is the number of victory points the player has.
	VictoryPoints() int64

	// Won is if the player won the game or not.
	Won() bool
}
