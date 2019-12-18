package anarchy

// FireDepartment is can put out fires completely.
type FireDepartment interface {
	// Parent interfaces
	Building

	// -- Attributes -- \\

	// FireExtinguished is the amount of fire removed from a building
	// when bribed to extinguish a building.
	FireExtinguished() int64

	// -- Methods -- \\

	// Extinguish bribes this FireDepartment to extinguish the some of
	// the fire in a building.
	Extinguish(Building) bool
}
