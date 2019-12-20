package newtonian

// Machine is a machine in the game. Used to refine ore.
type Machine interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// OreType is what type of ore the machine takes it. Also
	// determines the type of material it outputs. (redium or
	// blueium).
//
	// Literal Values: "redium" or "blueium"
	OreType() string

	// RefineInput is the amount of ore that needs to be inputted into
	// the machine for it to be worked.
	RefineInput() int64

	// RefineOutput is the amount of refined ore that is returned
	// after the machine has been fully worked.
	RefineOutput() int64

	// RefineTime is the number of times this machine needs to be
	// worked to refine ore.
	RefineTime() int64

	// Tile is the Tile this Machine is on.
	Tile() Tile

	// Worked is tracks how many times this machine has been worked.
	// (0 to refineTime).
	Worked() int64
}
