package stardash

import "joueur/base"

// Game is collect of the most of the rarest mineral orbiting around the
// sun and out-compete your competitor.
type Game interface {
	// Parent interfaces
	base.Game

	// -- Attributes -- \\

	// Bodies is all the celestial bodies in the game. The first two
	// are planets and the third is the sun. The fourth is the VP
	// asteroid. Everything else is normal asteroids.
	Bodies() []Body

	// CurrentPlayer is the player whose turn it is currently. That
	// player can send commands. Other players cannot.
	CurrentPlayer() Player

	// CurrentTurn is the current turn number, starting at 0 for the
	// first player's turn.
	CurrentTurn() int64

	// DashCost is the cost of dashing.
	DashCost() int64

	// DashDistance is the distance traveled each turn by dashing.
	DashDistance() int64

	// GameObjects is a mapping of every game object's ID to the
	// actual game object. Primarily used by the server and client to
	// easily refer to the game objects via ID.
	GameObjects() map[string]GameObject

	// GenariumValue is the value of every unit of genarium.
	GenariumValue() float64

	// Jobs is an array of all jobs. The first element is corvette,
	// second is missileboat, third is martyr, fourth is transport,
	// and fifth is miner.
	Jobs() []Job

	// LegendariumValue is the value of every unit of legendarium.
	LegendariumValue() float64

	// MaxAsteroid is the highest amount of material, that can be in a
	// asteroid.
	MaxAsteroid() int64

	// MaxTurns is the maximum number of turns before the game will
	// automatically end.
	MaxTurns() int64

	// MinAsteroid is the smallest amount of material, that can be in
	// a asteroid.
	MinAsteroid() int64

	// MiningSpeed is the rate at which miners grab minerals from
	// asteroids.
	MiningSpeed() int64

	// MythiciteAmount is the amount of mythicite that spawns at the
	// start of the game.
	MythiciteAmount() float64

	// OrbitsProtected is the number of orbit updates you cannot mine
	// the mithicite asteroid.
	OrbitsProtected() int64

	// OreRarityGenarium is the rarity modifier of the most common
	// ore. This controls how much spawns.
	OreRarityGenarium() float64

	// OreRarityLegendarium is the rarity modifier of the rarest ore.
	// This controls how much spawns.
	OreRarityLegendarium() float64

	// OreRarityRarium is the rarity modifier of the second rarest
	// ore. This controls how much spawns.
	OreRarityRarium() float64

	// PlanetEnergyCap is the amount of energy a planet can hold at
	// once.
	PlanetEnergyCap() int64

	// PlanetRechargeRate is the amount of energy the planets restore
	// each round.
	PlanetRechargeRate() int64

	// Players is array of all the players in the game.
	Players() []Player

	// ProjectileRadius is the standard size of ships.
	ProjectileRadius() int64

	// ProjectileSpeed is the amount of distance missiles travel
	// through space.
	ProjectileSpeed() int64

	// Projectiles is every projectile in the game.
	Projectiles() []Projectile

	// RariumValue is the value of every unit of rarium.
	RariumValue() float64

	// RegenerateRate is the regeneration rate of asteroids.
	RegenerateRate() float64

	// Session is a unique identifier for the game instance that is
	// being played.
	Session() string

	// ShipRadius is the standard size of ships.
	ShipRadius() int64

	// SizeX is the size of the map in the X direction.
	SizeX() int64

	// SizeY is the size of the map in the Y direction.
	SizeY() int64

	// TimeAddedPerTurn is the amount of time (in nano-seconds) added
	// after each player performs a turn.
	TimeAddedPerTurn() int64

	// TurnsToOrbit is the number of turns it takes for a asteroid to
	// orbit the sun. (Asteroids move after each players turn).
	TurnsToOrbit() int64

	// Units is every Unit in the game.
	Units() []Unit
}
