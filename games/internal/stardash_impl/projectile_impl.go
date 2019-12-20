package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/stardash"
)

// ProjectileImpl is the struct that implements the container for
// Projectile instances in Stardash.
type ProjectileImpl struct {
	GameObjectImpl

	energyImpl int64
	fuelImpl   int64
	ownerImpl  stardash.Player
	targetImpl stardash.Unit
	xImpl      float64
	yImpl      float64
}

// Energy returns the remaining health of the projectile.
func (projectileImpl *ProjectileImpl) Energy() int64 {
	return projectileImpl.energyImpl
}

// Fuel returns the amount of remaining distance the projectile can move.
func (projectileImpl *ProjectileImpl) Fuel() int64 {
	return projectileImpl.fuelImpl
}

// Owner returns the Player that owns and can control this Projectile.
//
// Value can be returned as a nil pointer.
func (projectileImpl *ProjectileImpl) Owner() stardash.Player {
	return projectileImpl.ownerImpl
}

// Target returns the unit that is being attacked by this projectile.
func (projectileImpl *ProjectileImpl) Target() stardash.Unit {
	return projectileImpl.targetImpl
}

// X returns the x value this projectile is on.
func (projectileImpl *ProjectileImpl) X() float64 {
	return projectileImpl.xImpl
}

// Y returns the y value this projectile is on.
func (projectileImpl *ProjectileImpl) Y() float64 {
	return projectileImpl.yImpl
}

// InitImplDefaults initializes safe defaults for all fields in Projectile.
func (projectileImpl *ProjectileImpl) InitImplDefaults() {
	projectileImpl.GameObjectImpl.InitImplDefaults()

	projectileImpl.energyImpl = 0
	projectileImpl.fuelImpl = 0
	projectileImpl.ownerImpl = nil
	projectileImpl.targetImpl = nil
	projectileImpl.xImpl = 0
	projectileImpl.yImpl = 0
}

// DeltaMerge merges the delta for a given attribute in Projectile.
func (projectileImpl *ProjectileImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := projectileImpl.GameObjectImpl.DeltaMerge(
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
	case "energy":
		projectileImpl.energyImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "fuel":
		projectileImpl.fuelImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "owner":
		projectileImpl.ownerImpl = stardashDeltaMerge.Player(delta)
		return true, nil
	case "target":
		projectileImpl.targetImpl = stardashDeltaMerge.Unit(delta)
		return true, nil
	case "x":
		projectileImpl.xImpl = stardashDeltaMerge.Float(delta)
		return true, nil
	case "y":
		projectileImpl.yImpl = stardashDeltaMerge.Float(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}
