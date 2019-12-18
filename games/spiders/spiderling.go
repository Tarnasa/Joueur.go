package spiders

// Spiderling is a Spider spawned by the BroodMother.
type Spiderling interface {
	// Parent interfaces
	Spider

	// -- Attributes -- \\

	// Busy is when empty string this Spiderling is not busy, and can
	// act. Otherwise a string representing what it is busy with, e.g.
	// 'Moving', 'Attacking'.
	// Literal Values: "", "Moving", "Attacking", "Strengthening",
	// "Weakening", "Cutting", or "Spitting"
	Busy() string

	// MovingOnWeb is the Web this Spiderling is using to move. Nil if
	// it is not moving.
	MovingOnWeb() Web

	// MovingToNest is the Nest this Spiderling is moving to. Nil if
	// it is not moving.
	MovingToNest() Nest

	// NumberOfCoworkers is the number of Spiderlings busy with the
	// same work this Spiderling is doing, speeding up the task.
	NumberOfCoworkers() int64

	// WorkRemaining is how much work needs to be done for this
	// Spiderling to finish being busy. See docs for the Work forumla.
	WorkRemaining() float64

	// -- Methods -- \\

	// Attack attacks another Spiderling
	Attack(Spiderling) bool

	// Move starts moving the Spiderling across a Web to another Nest.
	Move(Web) bool
}
