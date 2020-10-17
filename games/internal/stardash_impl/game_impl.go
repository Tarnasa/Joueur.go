package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/stardash"
)

// GameImpl is the struct that implements the container for Game instances
// in Stardash.
type GameImpl struct {
	base.GameImpl

	bodiesImpl               []stardash.Body
	currentPlayerImpl        stardash.Player
	currentTurnImpl          int64
	dashCostImpl             int64
	dashDistanceImpl         int64
	gameObjectsImpl          map[string]stardash.GameObject
	genariumValueImpl        float64
	jobsImpl                 []stardash.Job
	legendariumValueImpl     float64
	maxAsteroidImpl          int64
	maxTurnsImpl             int64
	minAsteroidImpl          int64
	miningSpeedImpl          int64
	mythiciteAmountImpl      float64
	orbitsProtectedImpl      int64
	oreRarityGenariumImpl    float64
	oreRarityLegendariumImpl float64
	oreRarityRariumImpl      float64
	planetEnergyCapImpl      int64
	planetRechargeRateImpl   int64
	playersImpl              []stardash.Player
	projectileRadiusImpl     int64
	projectileSpeedImpl      int64
	projectilesImpl          []stardash.Projectile
	rariumValueImpl          float64
	regenerateRateImpl       float64
	sessionImpl              string
	shipRadiusImpl           int64
	sizeXImpl                int64
	sizeYImpl                int64
	timeAddedPerTurnImpl     int64
	turnsToOrbitImpl         int64
	unitsImpl                []stardash.Unit
}

// Bodies returns all the celestial bodies in the game. The first two are
// planets and the third is the sun. The fourth is the VP asteroid.
// Everything else is normal asteroids.
func (gameImpl *GameImpl) Bodies() []stardash.Body {
	return gameImpl.bodiesImpl
}

// CurrentPlayer returns the player whose turn it is currently. That player
// can send commands. Other players cannot.
func (gameImpl *GameImpl) CurrentPlayer() stardash.Player {
	return gameImpl.currentPlayerImpl
}

// CurrentTurn returns the current turn number, starting at 0 for the first
// player's turn.
func (gameImpl *GameImpl) CurrentTurn() int64 {
	return gameImpl.currentTurnImpl
}

// DashCost returns the cost of dashing.
func (gameImpl *GameImpl) DashCost() int64 {
	return gameImpl.dashCostImpl
}

// DashDistance returns the distance traveled each turn by dashing.
func (gameImpl *GameImpl) DashDistance() int64 {
	return gameImpl.dashDistanceImpl
}

// GameObjects returns a mapping of every game object's ID to the actual
// game object. Primarily used by the server and client to easily refer to
// the game objects via ID.
func (gameImpl *GameImpl) GameObjects() map[string]stardash.GameObject {
	return gameImpl.gameObjectsImpl
}

// GenariumValue returns the value of every unit of genarium.
func (gameImpl *GameImpl) GenariumValue() float64 {
	return gameImpl.genariumValueImpl
}

// Jobs returns an array of all jobs. The first element is corvette, second
// is missileboat, third is martyr, fourth is transport, and fifth is
// miner.
func (gameImpl *GameImpl) Jobs() []stardash.Job {
	return gameImpl.jobsImpl
}

// LegendariumValue returns the value of every unit of legendarium.
func (gameImpl *GameImpl) LegendariumValue() float64 {
	return gameImpl.legendariumValueImpl
}

// MaxAsteroid returns the highest amount of material, that can be in a
// asteroid.
func (gameImpl *GameImpl) MaxAsteroid() int64 {
	return gameImpl.maxAsteroidImpl
}

// MaxTurns returns the maximum number of turns before the game will
// automatically end.
func (gameImpl *GameImpl) MaxTurns() int64 {
	return gameImpl.maxTurnsImpl
}

// MinAsteroid returns the smallest amount of material, that can be in a
// asteroid.
func (gameImpl *GameImpl) MinAsteroid() int64 {
	return gameImpl.minAsteroidImpl
}

// MiningSpeed returns the rate at which miners grab minerals from
// asteroids.
func (gameImpl *GameImpl) MiningSpeed() int64 {
	return gameImpl.miningSpeedImpl
}

// MythiciteAmount returns the amount of mythicite that spawns at the start
// of the game.
func (gameImpl *GameImpl) MythiciteAmount() float64 {
	return gameImpl.mythiciteAmountImpl
}

// OrbitsProtected returns the number of orbit updates you cannot mine the
// mithicite asteroid.
func (gameImpl *GameImpl) OrbitsProtected() int64 {
	return gameImpl.orbitsProtectedImpl
}

// OreRarityGenarium returns the rarity modifier of the most common ore.
// This controls how much spawns.
func (gameImpl *GameImpl) OreRarityGenarium() float64 {
	return gameImpl.oreRarityGenariumImpl
}

// OreRarityLegendarium returns the rarity modifier of the rarest ore. This
// controls how much spawns.
func (gameImpl *GameImpl) OreRarityLegendarium() float64 {
	return gameImpl.oreRarityLegendariumImpl
}

// OreRarityRarium returns the rarity modifier of the second rarest ore.
// This controls how much spawns.
func (gameImpl *GameImpl) OreRarityRarium() float64 {
	return gameImpl.oreRarityRariumImpl
}

// PlanetEnergyCap returns the amount of energy a planet can hold at once.
func (gameImpl *GameImpl) PlanetEnergyCap() int64 {
	return gameImpl.planetEnergyCapImpl
}

// PlanetRechargeRate returns the amount of energy the planets restore each
// round.
func (gameImpl *GameImpl) PlanetRechargeRate() int64 {
	return gameImpl.planetRechargeRateImpl
}

// Players returns array of all the players in the game.
func (gameImpl *GameImpl) Players() []stardash.Player {
	return gameImpl.playersImpl
}

// ProjectileRadius returns the standard size of ships.
func (gameImpl *GameImpl) ProjectileRadius() int64 {
	return gameImpl.projectileRadiusImpl
}

// ProjectileSpeed returns the amount of distance missiles travel through
// space.
func (gameImpl *GameImpl) ProjectileSpeed() int64 {
	return gameImpl.projectileSpeedImpl
}

// Projectiles returns every projectile in the game.
func (gameImpl *GameImpl) Projectiles() []stardash.Projectile {
	return gameImpl.projectilesImpl
}

// RariumValue returns the value of every unit of rarium.
func (gameImpl *GameImpl) RariumValue() float64 {
	return gameImpl.rariumValueImpl
}

// RegenerateRate returns the regeneration rate of asteroids.
func (gameImpl *GameImpl) RegenerateRate() float64 {
	return gameImpl.regenerateRateImpl
}

// Session returns a unique identifier for the game instance that is being
// played.
func (gameImpl *GameImpl) Session() string {
	return gameImpl.sessionImpl
}

// ShipRadius returns the standard size of ships.
func (gameImpl *GameImpl) ShipRadius() int64 {
	return gameImpl.shipRadiusImpl
}

// SizeX returns the size of the map in the X direction.
func (gameImpl *GameImpl) SizeX() int64 {
	return gameImpl.sizeXImpl
}

// SizeY returns the size of the map in the Y direction.
func (gameImpl *GameImpl) SizeY() int64 {
	return gameImpl.sizeYImpl
}

// TimeAddedPerTurn returns the amount of time (in nano-seconds) added
// after each player performs a turn.
func (gameImpl *GameImpl) TimeAddedPerTurn() int64 {
	return gameImpl.timeAddedPerTurnImpl
}

// TurnsToOrbit returns the number of turns it takes for a asteroid to
// orbit the sun. (Asteroids move after each players turn).
func (gameImpl *GameImpl) TurnsToOrbit() int64 {
	return gameImpl.turnsToOrbitImpl
}

// Units returns every Unit in the game.
func (gameImpl *GameImpl) Units() []stardash.Unit {
	return gameImpl.unitsImpl
}

// InitImplDefaults initializes safe defaults for all fields in Game.
func (gameImpl *GameImpl) InitImplDefaults() {
	gameImpl.GameImpl.InitImplDefaults()

	gameImpl.bodiesImpl = []stardash.Body{}
	gameImpl.currentPlayerImpl = nil
	gameImpl.currentTurnImpl = 0
	gameImpl.dashCostImpl = 0
	gameImpl.dashDistanceImpl = 0
	gameImpl.gameObjectsImpl = make(map[string]stardash.GameObject)
	gameImpl.genariumValueImpl = 0
	gameImpl.jobsImpl = []stardash.Job{}
	gameImpl.legendariumValueImpl = 0
	gameImpl.maxAsteroidImpl = 0
	gameImpl.maxTurnsImpl = 0
	gameImpl.minAsteroidImpl = 0
	gameImpl.miningSpeedImpl = 0
	gameImpl.mythiciteAmountImpl = 0
	gameImpl.orbitsProtectedImpl = 0
	gameImpl.oreRarityGenariumImpl = 0
	gameImpl.oreRarityLegendariumImpl = 0
	gameImpl.oreRarityRariumImpl = 0
	gameImpl.planetEnergyCapImpl = 0
	gameImpl.planetRechargeRateImpl = 0
	gameImpl.playersImpl = []stardash.Player{}
	gameImpl.projectileRadiusImpl = 0
	gameImpl.projectileSpeedImpl = 0
	gameImpl.projectilesImpl = []stardash.Projectile{}
	gameImpl.rariumValueImpl = 0
	gameImpl.regenerateRateImpl = 0
	gameImpl.sessionImpl = ""
	gameImpl.shipRadiusImpl = 0
	gameImpl.sizeXImpl = 0
	gameImpl.sizeYImpl = 0
	gameImpl.timeAddedPerTurnImpl = 0
	gameImpl.turnsToOrbitImpl = 0
	gameImpl.unitsImpl = []stardash.Unit{}
}

// DeltaMerge merges the delta for a given attribute in Game.
func (gameImpl *GameImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := gameImpl.GameImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	stardashDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'stardash.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "bodies":
		gameImpl.bodiesImpl = stardashDeltaMerge.ArrayOfBody(&gameImpl.bodiesImpl, delta)
		return true, nil
	case "currentPlayer":
		gameImpl.currentPlayerImpl = stardashDeltaMerge.Player(delta)
		return true, nil
	case "currentTurn":
		gameImpl.currentTurnImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "dashCost":
		gameImpl.dashCostImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "dashDistance":
		gameImpl.dashDistanceImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "gameObjects":
		gameImpl.gameObjectsImpl = stardashDeltaMerge.MapOfStringToGameObject(&gameImpl.gameObjectsImpl, delta)
		return true, nil
	case "genariumValue":
		gameImpl.genariumValueImpl = stardashDeltaMerge.Float(delta)
		return true, nil
	case "jobs":
		gameImpl.jobsImpl = stardashDeltaMerge.ArrayOfJob(&gameImpl.jobsImpl, delta)
		return true, nil
	case "legendariumValue":
		gameImpl.legendariumValueImpl = stardashDeltaMerge.Float(delta)
		return true, nil
	case "maxAsteroid":
		gameImpl.maxAsteroidImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "maxTurns":
		gameImpl.maxTurnsImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "minAsteroid":
		gameImpl.minAsteroidImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "miningSpeed":
		gameImpl.miningSpeedImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "mythiciteAmount":
		gameImpl.mythiciteAmountImpl = stardashDeltaMerge.Float(delta)
		return true, nil
	case "orbitsProtected":
		gameImpl.orbitsProtectedImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "oreRarityGenarium":
		gameImpl.oreRarityGenariumImpl = stardashDeltaMerge.Float(delta)
		return true, nil
	case "oreRarityLegendarium":
		gameImpl.oreRarityLegendariumImpl = stardashDeltaMerge.Float(delta)
		return true, nil
	case "oreRarityRarium":
		gameImpl.oreRarityRariumImpl = stardashDeltaMerge.Float(delta)
		return true, nil
	case "planetEnergyCap":
		gameImpl.planetEnergyCapImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "planetRechargeRate":
		gameImpl.planetRechargeRateImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "players":
		gameImpl.playersImpl = stardashDeltaMerge.ArrayOfPlayer(&gameImpl.playersImpl, delta)
		return true, nil
	case "projectileRadius":
		gameImpl.projectileRadiusImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "projectileSpeed":
		gameImpl.projectileSpeedImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "projectiles":
		gameImpl.projectilesImpl = stardashDeltaMerge.ArrayOfProjectile(&gameImpl.projectilesImpl, delta)
		return true, nil
	case "rariumValue":
		gameImpl.rariumValueImpl = stardashDeltaMerge.Float(delta)
		return true, nil
	case "regenerateRate":
		gameImpl.regenerateRateImpl = stardashDeltaMerge.Float(delta)
		return true, nil
	case "session":
		gameImpl.sessionImpl = stardashDeltaMerge.String(delta)
		return true, nil
	case "shipRadius":
		gameImpl.shipRadiusImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "sizeX":
		gameImpl.sizeXImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "sizeY":
		gameImpl.sizeYImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "timeAddedPerTurn":
		gameImpl.timeAddedPerTurnImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "turnsToOrbit":
		gameImpl.turnsToOrbitImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "units":
		gameImpl.unitsImpl = stardashDeltaMerge.ArrayOfUnit(&gameImpl.unitsImpl, delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}
