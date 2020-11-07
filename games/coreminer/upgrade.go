package coreminer

// Upgrade is information about a Miner's Upgrade module.
type Upgrade interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// CargoCapacity is the amount of cargo capacity this Upgrade has.
	CargoCapacity() int64

	// Health is the maximum amount of health this Upgrade has.
	Health() int64

	// MiningPower is the amount of mining power this Upgrade has per
	// turn.
	MiningPower() int64

	// Moves is the number of moves this Upgrade can make per turn.
	Moves() int64

	// Title is the Upgrade title.
	Title() string
}
