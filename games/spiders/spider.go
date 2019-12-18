package spiders

// Spider is a Spider in the game. The most basic unit.
type Spider interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// IsDead is if this Spider is dead and has been removed from the
	// game.
	IsDead() bool

	// Nest is the Nest that this Spider is currently on. Nil when
	// moving on a Web.
	Nest() Nest

	// Owner is the Player that owns this Spider, and can command it.
	Owner() Player
}
