package spiders

// Spitter is a Spiderling that creates and spits new Webs from the Nest it
// is on to another Nest, connecting them.
type Spitter interface {
	// Parent interfaces
	Spiderling

	// -- Attributes -- \\

	// SpittingWebToNest is the Nest that this Spitter is creating a
	// Web to spit at, thus connecting them. Nil if not spitting.
	//
	// Value can be returned as a nil pointer.
	SpittingWebToNest() Nest

	// -- Methods -- \\

	// Spit creates and spits a new Web from the Nest the Spitter is
	// on to another Nest, connecting them.
	Spit(Nest) bool
}
