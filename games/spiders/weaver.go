package spiders

// Weaver is a Spiderling that can alter existing Webs by weaving to add or
// remove silk from the Webs, thus altering its strength.
type Weaver interface {
	// Parent interfaces
	Spiderling

	// -- Attributes -- \\

	// StrengtheningWeb is the Web that this Weaver is strengthening.
	// Nil if not strengthening.
	StrengtheningWeb() Web

	// WeakeningWeb is the Web that this Weaver is weakening. Nil if
	// not weakening.
	WeakeningWeb() Web

	// -- Methods -- \\

	// Strengthen weaves more silk into an existing Web to strengthen
	// it.
	Strengthen(Web) bool

	// Weaken weaves more silk into an existing Web to strengthen it.
	Weaken(Web) bool
}
