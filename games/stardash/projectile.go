package stardash

// Projectile is tracks any projectiles moving through space.
type Projectile interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Energy is the remaining health of the projectile.
	Energy() int64

	// Fuel is the amount of remaining distance the projectile can
	// move.
	Fuel() int64

	// Owner is the Player that owns and can control this Projectile.
	//
	// Value can be returned as a nil pointer.
	Owner() Player

	// Target is the unit that is being attacked by this projectile.
	Target() Unit

	// X is the x value this projectile is on.
	X() float64

	// Y is the y value this projectile is on.
	Y() float64
}
