package pirates

// Unit is a unit group in the game. This may consist of a ship and any
// number of crew.
type Unit interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Acted is whether this Unit has performed its action this turn.
	Acted() bool

	// Crew is how many crew are on this Tile. This number will always
	// be <= crewHealth.
	Crew() int64

	// CrewHealth is how much total health the crew on this Tile have.
	CrewHealth() int64

	// Gold is how much gold this Unit is carrying.
	Gold() int64

	// Moves is how many more times this Unit may move this turn.
	Moves() int64

	// Owner is the Player that owns and can control this Unit, or nil
	// if the Unit is neutral.
	Owner() Player

	// Path is (Merchants only) The path this Unit will follow. The
	// first element is the Tile this Unit will move to next.
	Path() []Tile

	// ShipHealth is if a ship is on this Tile, how much health it has
	// remaining. 0 for no ship.
	ShipHealth() int64

	// StunTurns is (Merchants only) The number of turns this merchant
	// ship won't be able to move. They will still attack. Merchant
	// ships are stunned when they're attacked.
	StunTurns() int64

	// TargetPort is (Merchants only) The Port this Unit is moving to.
	TargetPort() Port

	// Tile is the Tile this Unit is on.
	Tile() Tile

	// -- Methods -- \\

	// Attack attacks either the 'crew' or 'ship' on a Tile in range.
	Attack(Tile, string) bool

	// Bury buries gold on this Unit's Tile. Gold must be a certain
	// distance away for it to get interest
	// (Game.minInterestDistance).
	Bury(int64) bool

	// Deposit puts gold into an adjacent Port. If that Port is the
	// Player's port, the gold is added to that Player. If that Port
	// is owned by merchants, it adds to that Port's investment.
	Deposit(int64) bool

	// Dig digs up gold on this Unit's Tile.
	Dig(int64) bool

	// Move moves this Unit from its current Tile to an adjacent Tile.
	// If this Unit merges with another one, the other Unit will be
	// destroyed and its tile will be set to nil. Make sure to check
	// that your Unit's tile is not nil before doing things with it.
	Move(Tile) bool

	// Rest regenerates this Unit's health. Must be used in range of a
	// port.
	Rest() bool

	// Split moves a number of crew from this Unit to the given Tile.
	// This will consume a move from those crew.
	Split(Tile, int64, int64) bool

	// Withdraw takes gold from the Player. You can only withdraw from
	// your own Port.
	Withdraw(int64) bool
}
