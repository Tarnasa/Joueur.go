package anarchy

// Warehouse is a typical abandoned warehouse that anarchists hang out in
// and can be bribed to burn down Buildings.
type Warehouse interface {
	// Parent interfaces
	Building

	// -- Attributes -- \\

	// Exposure is how exposed the anarchists in this warehouse are to
	// PoliceDepartments. Raises when bribed to ignite buildings, and
	// drops each turn if not bribed.
	Exposure() int64

	// FireAdded is the amount of fire added to buildings when bribed
	// to ignite a building. Headquarters add more fire than normal
	// Warehouses.
	FireAdded() int64

	// -- Methods -- \\

	// Ignite bribes the Warehouse to light a Building on fire. This
	// adds this building's fireAdded to their fire, and then this
	// building's exposure is increased based on the Manhattan
	// distance between the two buildings.
	Ignite(Building) int64
}
