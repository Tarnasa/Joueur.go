package pirates

// Port is a port on a Tile.
type Port interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Gold is for players, how much more gold this Port can spend
	// this turn. For merchants, how much gold this Port has
	// accumulated (it will spawn a ship when the Port can afford
	// one).
	Gold() int64

	// Investment is (Merchants only) How much gold was invested into
	// this Port. Investment determines the strength and value of the
	// next ship.
	Investment() int64

	// Owner is the owner of this Port, or nil if owned by merchants.
	//
	// Value can be returned as a nil pointer.
	Owner() Player

	// Tile is the Tile this Port is on.
	Tile() Tile

	// -- Methods -- \\

	// Spawn spawn a Unit on this port.
	Spawn(string) bool
}
