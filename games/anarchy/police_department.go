package anarchy

// PoliceDepartment is used to keep cities under control and raid
// Warehouses.
type PoliceDepartment interface {
	// Parent interfaces
	Building

	// -- Methods -- \\

	// Raid bribe the police to raid a Warehouse, dealing damage equal
	// based on the Warehouse's current exposure, and then resetting
	// it to 0.
	Raid(Warehouse) int64
}
