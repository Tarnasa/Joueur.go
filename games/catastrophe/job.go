package catastrophe

// Job is information about a Unit's job.
type Job interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// ActionCost is the amount of energy this Job normally uses to
	// perform its actions.
	ActionCost() float64

	// CarryLimit is how many combined resources a Unit with this Job
	// can hold at once.
	CarryLimit() int64

	// Moves is the number of moves this Job can make per turn.
	Moves() int64

	// RegenRate is the amount of energy normally regenerated when
	// resting at a shelter.
	RegenRate() float64

	// Title is the Job title.
	// Literal Values: "fresh human", "cat overlord", "soldier",
	// "gatherer", "builder", or "missionary"
	Title() string

	// Upkeep is the amount of food per turn this Unit consumes. If
	// there isn't enough food for every Unit, all Units become
	// starved and do not consume food.
	Upkeep() int64
}
