package anarchy

// Player is a player in this game. Every AI controls one player.
type Player interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// BribesRemaining is how many bribes this player has remaining to
	// use during their turn. Each action a Building does costs 1
	// bribe. Any unused bribes are lost at the end of the player's
	// turn.
	BribesRemaining() int64

	// Buildings is all the buildings owned by this player.
	Buildings() []Building

	// ClientType is what type of client this is, e.g. 'Python',
	// 'JavaScript', or some other language. For potential data mining
	// purposes.
	ClientType() string

	// FireDepartments is all the FireDepartments owned by this
	// player.
	FireDepartments() []FireDepartment

	// Headquarters is the Warehouse that serves as this player's
	// headquarters and has extra health. If this gets destroyed they
	// lose.
	Headquarters() Warehouse

	// Lost is if the player lost the game or not.
	Lost() bool

	// Name is the name of the player.
	Name() string

	// Opponent is this player's opponent in the game.
	Opponent() Player

	// PoliceDepartments is all the PoliceDepartments owned by this
	// player.
	PoliceDepartments() []PoliceDepartment

	// ReasonLost is the reason why the player lost the game.
	ReasonLost() string

	// ReasonWon is the reason why the player won the game.
	ReasonWon() string

	// TimeRemaining is the amount of time (in ns) remaining for this
	// AI to send commands.
	TimeRemaining() float64

	// Warehouses is all the warehouses owned by this player. Includes
	// the Headquarters.
	Warehouses() []Warehouse

	// WeatherStations is all the WeatherStations owned by this
	// player.
	WeatherStations() []WeatherStation

	// Won is if the player won the game or not.
	Won() bool
}
