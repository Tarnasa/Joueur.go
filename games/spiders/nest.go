package spiders

// Nest is a location (node) connected to other Nests via Webs (edges) in
// the game that Spiders can converge on, regardless of owner.
type Nest interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// ControllingPlayer is the Player that 'controls' this Nest as
	// they have the most Spiders on this nest.
	//
	// Value can be returned as a nil pointer.
	ControllingPlayer() Player

	// Spiders is all the Spiders currently located on this Nest.
	Spiders() []Spider

	// Webs is webs that connect to this Nest.
	Webs() []Web

	// X is the X coordinate of the Nest. Used for distance
	// calculations.
	X() int64

	// Y is the Y coordinate of the Nest. Used for distance
	// calculations.
	Y() int64
}
