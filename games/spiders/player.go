package spiders

// Player is a player in this game. Every AI controls one player.
type Player interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// BroodMother is this player's BroodMother. If it dies they lose
	// the game.
	BroodMother() BroodMother

	// ClientType is what type of client this is, e.g. 'Python',
	// 'JavaScript', or some other language. For potential data mining
	// purposes.
	ClientType() string

	// Lost is if the player lost the game or not.
	Lost() bool

	// MaxSpiderlings is the max number of Spiderlings players can
	// spawn.
	MaxSpiderlings() int64

	// Name is the name of the player.
	Name() string

	// NumberOfNestsControlled is the number of nests this player
	// controls.
	NumberOfNestsControlled() int64

	// Opponent is this player's opponent in the game.
	Opponent() Player

	// ReasonLost is the reason why the player lost the game.
	ReasonLost() string

	// ReasonWon is the reason why the player won the game.
	ReasonWon() string

	// Spiders is all the Spiders owned by this player.
	Spiders() []Spider

	// TimeRemaining is the amount of time (in ns) remaining for this
	// AI to send commands.
	TimeRemaining() float64

	// Won is if the player won the game or not.
	Won() bool
}
