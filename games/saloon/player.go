package saloon

// Player is a player in this game. Every AI controls one player.
type Player interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// ClientType is what type of client this is, e.g. 'Python',
	// 'JavaScript', or some other language. For potential data mining
	// purposes.
	ClientType() string

	// Cowboys is every Cowboy owned by this Player.
	Cowboys() []Cowboy

	// Kills is how many enemy Cowboys this player's team has killed.
	Kills() int64

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

	// Rowdiness is how rowdy their team is. When it gets too high
	// their team takes a collective siesta.
	Rowdiness() int64

	// Score is how many times their team has played a piano.
	Score() int64

	// Siesta is 0 when not having a team siesta. When greater than 0
	// represents how many turns left for the team siesta to complete.
	Siesta() int64

	// TimeRemaining is the amount of time (in ns) remaining for this
	// AI to send commands.
	TimeRemaining() float64

	// Won is if the player won the game or not.
	Won() bool

	// YoungGun is the YoungGun this Player uses to call in new
	// Cowboys.
	YoungGun() YoungGun
}
