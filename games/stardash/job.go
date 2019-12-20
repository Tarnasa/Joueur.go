package stardash

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

	// Energy is the amount of starting health this Job has.
	Energy() int64

	// Moves is the distance this job can move per turn.
	Moves() int64

	// Range is the distance at which this job can effect things.
	Range() int64

	// Shield is the reserve the martyr use to protect allies.
	Shield() int64

	// Title is the Job title. 'corvette', 'missileboat', 'martyr',
	// 'transport', or 'miner'. (in this order from 0-4).
//
	// Literal Values: "corvette", "missileboat", "martyr",
	// "transport", or "miner"
	Title() string

	// UnitCost is how much money it costs to spawn a unit.
	UnitCost() int64
}
