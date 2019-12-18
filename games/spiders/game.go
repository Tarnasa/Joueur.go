package spiders

import "joueur/base"

// Game is there's an infestation of enemy spiders challenging your queen
// broodmother spider! Protect her and attack the other broodmother in this
// turn based, node based, game.
type Game interface {
	// Parent interfaces
	base.Game

	// -- Attributes -- \\

	// CurrentPlayer is the player whose turn it is currently. That
	// player can send commands. Other players cannot.
	CurrentPlayer() Player

	// CurrentTurn is the current turn number, starting at 0 for the
	// first player's turn.
	CurrentTurn() int64

	// CutSpeed is the speed at which Cutters work to do cut Webs.
	CutSpeed() int64

	// EggsScalar is constant used to calculate how many eggs
	// BroodMothers get on their owner's turns.
	EggsScalar() float64

	// GameObjects is a mapping of every game object's ID to the
	// actual game object. Primarily used by the server and client to
	// easily refer to the game objects via ID.
	GameObjects() map[string]GameObject

	// InitialWebStrength is the starting strength for Webs.
	InitialWebStrength() int64

	// MaxTurns is the maximum number of turns before the game will
	// automatically end.
	MaxTurns() int64

	// MaxWebStrength is the maximum strength a web can be
	// strengthened to.
	MaxWebStrength() int64

	// MovementSpeed is the speed at which Spiderlings move on Webs.
	MovementSpeed() int64

	// Nests is every Nest in the game.
	Nests() []Nest

	// Players is array of all the players in the game.
	Players() []Player

	// Session is a unique identifier for the game instance that is
	// being played.
	Session() string

	// SpitSpeed is the speed at which Spitters work to spit new Webs.
	SpitSpeed() int64

	// TimeAddedPerTurn is the amount of time (in nano-seconds) added
	// after each player performs a turn.
	TimeAddedPerTurn() int64

	// WeavePower is how much web strength is added or removed from
	// Webs when they are weaved.
	WeavePower() int64

	// WeaveSpeed is the speed at which Weavers work to do strengthens
	// and weakens on Webs.
	WeaveSpeed() int64

	// Webs is every Web in the game.
	Webs() []Web
}
