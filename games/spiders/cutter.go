package spiders

// Cutter is a Spiderling that can cut existing Webs.
type Cutter interface {
	// Parent interfaces
	Spiderling

	// -- Attributes -- \\

	// CuttingWeb is the Web that this Cutter is trying to cut. Nil if
	// not cutting.
	//
	// Value can be returned as a nil pointer.
	CuttingWeb() Web

	// -- Methods -- \\

	// Cut cuts a web, destroying it, and any Spiderlings on it.
	Cut(Web) bool
}
