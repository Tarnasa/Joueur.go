package stumped

// Job is information about a beaver's job.
type Job interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Actions is the number of actions this Job can make per turn.
	Actions() int64

	// CarryLimit is how many combined resources a beaver with this
	// Job can hold at once.
	CarryLimit() int64

	// Chopping is scalar for how many branches this Job harvests at
	// once.
	Chopping() int64

	// Cost is how much food this Job costs to recruit.
	Cost() int64

	// Damage is the amount of damage this Job does per attack.
	Damage() int64

	// DistractionPower is how many turns a beaver attacked by this
	// Job is distracted by.
	DistractionPower() int64

	// Health is the amount of starting health this Job has.
	Health() int64

	// Moves is the number of moves this Job can make per turn.
	Moves() int64

	// Munching is scalar for how much food this Job harvests at once.
	Munching() int64

	// Title is the Job title.
	Title() string

	// -- Methods -- \\

	// Recruit recruits a Beaver of this Job to a lodge.
	Recruit(Tile) Beaver
}
