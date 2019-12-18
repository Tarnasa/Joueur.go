package newtonian

// Job is information about a unit's job.
type Job interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// CarryLimit is how many combined resources a unit with this Job
	// can hold at once.
	CarryLimit() int64

	// Damage is the amount of damage this Job does per attack.
	Damage() int64

	// Health is the amount of starting health this Job has.
	Health() int64

	// Moves is the number of moves this Job can make per turn.
	Moves() int64

	// Title is the Job title. 'intern', 'manager', or 'physicist'.
	// Literal Values: "intern", "manager", or "physicist"
	Title() string
}
