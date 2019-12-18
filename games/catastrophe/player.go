package catastrophe

// Player is a player in this game. Every AI controls one player.
type Player interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Cat is the overlord cat Unit owned by this Player.
	Cat() Unit

	// ClientType is what type of client this is, e.g. 'Python',
	// 'JavaScript', or some other language. For potential data mining
	// purposes.
	ClientType() string

	// Food is the amount of food owned by this player.
	Food() int64

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

	// Structures is every Structure owned by this Player.
	Structures() []Structure

	// TimeRemaining is the amount of time (in ns) remaining for this
	// AI to send commands.
	TimeRemaining() float64

	// Units is every Unit owned by this Player.
	Units() []Unit

	// Upkeep is the total upkeep of every Unit owned by this Player.
	// If there isn't enough food for every Unit, all Units become
	// starved and do not consume food.
	Upkeep() int64

	// Won is if the player won the game or not.
	Won() bool
}
