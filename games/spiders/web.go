package spiders

// Web is a connection (edge) to a Nest (node) in the game that Spiders can
// converge on (regardless of owner). Spiders can travel in either
// direction on Webs.
type Web interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Length is how long this Web is, i.e., the distance between its
	// nestA and nestB.
	Length() float64

	// Load is how much weight this Web currently has on it, which is
	// the sum of all its Spiderlings weight.
	Load() int64

	// NestA is the first Nest this Web is connected to.
//
	// Value can be returned as a nil pointer.
	NestA() Nest

	// NestB is the second Nest this Web is connected to.
//
	// Value can be returned as a nil pointer.
	NestB() Nest

	// Spiderlings is all the Spiderlings currently moving along this
	// Web.
	Spiderlings() []Spiderling

	// Strength is how much weight this Web can take before snapping
	// and destroying itself and all the Spiders on it.
	Strength() int64
}
